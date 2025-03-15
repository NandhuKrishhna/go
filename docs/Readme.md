# Introduction to Go

**Go** is an open-source programming language designed for **simplicity**, **efficiency**, and **concurrency** (handling multiple tasks at once). It’s **statically typed** (meaning you declare variable types upfront) and **compiled** (turns your code into a single executable file). Think of it as a modern twist on languages like **C**, but way easier to read and write.

**What is a Package in Go?**

In Go, a package is like a container or a namespace that groups related code together. Think of it as a folder in your brain where you organize functions, variables, and types that belong together. Packages are Go’s way of:

- **Organizing code**: Keeping things modular and reusable.
- **Avoiding name clashes**: Two functions with the same name can coexist if they’re in different packages.
- **Encouraging collaboration**: You can import packages written by others (or yourself) to use their functionality.

Every Go file must belong to a package—it’s not optional. The first line of your file (e.g., `package main`) declares which package this file is part of.