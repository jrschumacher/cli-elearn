---
title: My test lesson

tags:
  - test
  - lesson
  - markdown

steps:
  - validator: '^\s*ls\s*$'
    exec: 'exec ls'
---

# Learning ls

In this lesson, we will learn how to use the `ls` command. The `ls` command
is used to list the contents of a directory.

For each operand that names a file of a type other than directory, ls
displays its name as well as any requested, associated information.  For
each operand that names a file of type directory, ls displays the names
of files contained within that directory, as well as any requested,
associated information.

If no operands are given, the contents of the current directory are
displayed.  If more than one operand is given, non-directory operands are
displayed first; directory and non-directory operands are sorted
separately and in lexicographical order.

## Getting the feel of it

Let's start by listing the contents of the current directory. To do this, type:

```bash
ls
```

## Step 2
