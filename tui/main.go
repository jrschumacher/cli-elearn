package tui

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/adrg/frontmatter"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"gopkg.in/yaml.v2"
)

const (
	termEntityUser = iota
	termEntityBot
	termEntityClass
)

const (
	viewportTerminal = iota
	viewportReader
)

type Lesson struct {
	Title string `yaml:"title"`
	Steps []struct {
		Title     string `yaml:"title"`
		Validator string `yaml:"validator"`
		Exec      string `yaml:"exec"`
	} `yaml:"steps"`
}

type (
	errMsg error

	terminalMessage struct {
		termEntity int
		msg        string
	}

	model struct {
		ready bool

		height int
		width  int

		err error

		instructions string
		lesson       Lesson
		currentStep  int

		input            textarea.Model
		terminal         viewport.Model
		terminalMessages []terminalMessage

		reader          viewport.Model
		readerCursor    int
		currentViewport int

		showHelper bool
	}
)

const (
	ZREADER = "reader"
	ZINPUT  = "input"
	ZTERM   = "term"
)

func initialModel() model {
	input := textarea.New()
	input.Placeholder = "Type here..."
	input.Focus()
	input.Prompt = "> "
	input.CharLimit = 280
	input.SetWidth(30)
	input.SetHeight(1)
	input.FocusedStyle.CursorLine = lipgloss.NewStyle()
	input.ShowLineNumbers = false
	input.KeyMap.InsertNewline.SetEnabled(false)

	terminal := viewport.New(30, 5)
	terminal.SetContent("Hello, world!")

	fn := "lessons/test.md"
	f, err := os.OpenFile(fn, os.O_RDONLY, 0644)
	if err != nil {
		slog.Error("Could not read file", slog.String("file", fn), slog.String("error", err.Error()))
		panic("Error reading file")
	}

	// parse front matter
	formats := []*frontmatter.Format{
		frontmatter.NewFormat("---", "---", yaml.Unmarshal),
	}
	lesson := Lesson{}
	rest, err := frontmatter.Parse(f, &lesson, formats...)
	if err != nil {
		slog.Error("Could not parse front matter", slog.String("file", fn), slog.String("error", err.Error()))
		panic("Error parsing front matter")
	}
	instructions := string(rest)

	return model{
		input: input,

		terminal: terminal,
		terminalMessages: []terminalMessage{
			{termEntityClass, "Hello, world!"},
		},

		lesson:       lesson,
		instructions: instructions,
	}
}

func (m model) Init() tea.Cmd {

	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.input, cmd = m.input.Update(msg)
	cmds = append(cmds, cmd)

	m.terminal, cmd = m.terminal.Update(msg)
	cmds = append(cmds, cmd)

	m.reader, cmd = m.reader.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			m = m.termProcess()
		}

	case tea.MouseMsg:
		if tea.MouseButton(msg.Button) == tea.MouseButtonLeft {
			if zone.Get(ZREADER).InBounds(msg) {
				m.currentViewport = viewportReader
				m.reader.MouseWheelEnabled = true
				m.terminal.MouseWheelEnabled = false
			} else if zone.Get(ZTERM).InBounds(msg) {
				m.currentViewport = viewportTerminal
				m.terminal.MouseWheelEnabled = true
				m.reader.MouseWheelEnabled = false
			}
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width

		verticalMarginHeight := 5
		headerHeight := 0
		useHighPerformanceRenderer := false
		if !m.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			m.reader = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.reader.KeyMap = viewport.KeyMap{
				Up: key.NewBinding(
					key.WithKeys("up", "ctrl+k"),
					key.WithHelp("↑/ctrl+k", "up"),
				),
				Down: key.NewBinding(
					key.WithKeys("down", "ctrl+j"),
					key.WithHelp("↓/ctrl+j", "down"),
				),
			}
			m.reader.YPosition = m.height - verticalMarginHeight
			m.reader.HighPerformanceRendering = useHighPerformanceRenderer

			// stylize with glamour and render
			r, _ := glamour.NewTermRenderer(
				// detect background color and pick either the default dark or light theme
				glamour.WithAutoStyle(),
				// wrap output at specific width (default is 80)
				glamour.WithWordWrap(m.width/2-6),
			)
			md, err := r.Render(m.instructions)
			if err != nil {
				md = m.instructions
			}
			m.reader.SetContent(md)
			fmt.Print(md)

			m.ready = true

			// This is only necessary for high performance rendering, which in
			// most cases you won't need.
			//
			// Render the viewport one line below the header.
			m.reader.YPosition = headerHeight + 1
		} else {
			m.reader.Width = msg.Width
			m.reader.Height = msg.Height - verticalMarginHeight
		}

		if useHighPerformanceRenderer {
			// Render (or re-render) the whole viewport. Necessary both to
			// initialize the viewport and when the window is resized.
			//
			// This is needed for high-performance rendering only.
			cmds = append(cmds, viewport.Sync(m.reader))
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	// Handle keyboard and mouse events in the viewport
	m.reader, cmd = m.reader.Update(msg)
	cmds = append(cmds, cmd)

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, tea.Batch(cmds...)
}

func (m model) View() string {

	borderDim := lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FFFFFF"}
	borderBright := lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#000000"}

	box := lipgloss.NewStyle().
		Width(m.width/2-4).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderDim).
		Padding(1, 1).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)

	boxFocus := box.Copy().BorderForeground(borderBright)

	helper := ""
	if m.showHelper {
		helper = box.Render("🤖 Howdy! Can I help you")
	}

	reader := "\n Loading..."
	if m.ready {
		reader = m.reader.View()
	}

	readerPane := box
	termPane := box
	if m.currentViewport == viewportReader {
		readerPane = boxFocus
	}
	if m.currentViewport == viewportTerminal {
		termPane = boxFocus
	}

	dialog := lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.JoinVertical(
			lipgloss.Top,
			zone.Mark(ZTERM, termPane.Render(m.terminal.View())),
			zone.Mark(ZINPUT, box.Render(m.input.View())),
		),
		lipgloss.JoinVertical(
			lipgloss.Top,
			zone.Mark(ZREADER, readerPane.Render(reader)),
			helper,
		),
	)

	return zone.Scan(dialog)
}

func Start() error {
	zone.NewGlobal()
	p := tea.NewProgram(initialModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}
