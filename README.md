# KSC (Known Shell Commands) 🚀

KSC is a Terminal User Interface (TUI) executable written in Go, designed to manage and execute frequently used shell commands efficiently. It allows users to store, search, and execute commands, including those with arguments, through an intuitive and interactive interface. 🖥️✨

## Features 🌟

- **Command Storage**: Add and manage frequently used shell commands. 📦
- **Search Functionality**: Quickly find stored commands using a search interface. 🔍
- **Argument Placeholders**: Store commands with placeholders (e.g., `ping <address>`) and prompt for arguments before execution. 📝➡️🔧
- **Interactive UI**: Beautiful and user-friendly terminal interface for seamless command management and execution. 🖼️🎨

## Installation 🛠️

1. **Clone the repository**:
    ```sh
    git clone https://github.com/yourusername/ksc.git
    cd ksc
    ```
    📂📥

2. **Install dependencies**:
    ```sh
    go mod tidy
    ```
    📦📋

3. **Build the executable**:
    ```sh
    go build -o ksc ./cmd/cli/*
    ```
    🏗️🚀

## Usage 🚀

### Adding Commands ➕

You can add commands to KSC, which can include placeholders for arguments. 📝

Example:
```sh
ksc add "ping <address>" "ping to specific address"
```
✨

### Searching and Selecting Commands 🔍

When you start KSC, you will see a list of added commands. Use the search functionality to filter through the list. 🔎

```sh
ksc
```
📜

### Executing Commands with Arguments 🛠️

If a command includes placeholders, KSC will prompt you to fill in the required arguments before execution. 📝➡️

Example:
1. Start KSC:
    ```sh
    ksc
    ```
    ▶️

2. Search and select the command `ping <address>`. 🔍📝

3. Enter the required argument when prompted:
    ```sh
    Enter value for <address>: 8.8.8.8
    ```
    🖊️➡️

KSC will then execute the command `ping 8.8.8.8`. 🏁🚀

## Example

Below is an example session using KSC: 🎬

```sh
# Add a command with an argument placeholder
ksc add "ping <address>" "ping to specific address"

# Start KSC
ksc

# Search and select the command
# (User types "ping" in the search box and selects "ping <address>")

# Prompt for the argument value
Enter value for <address>: 8.8.8.8

# Command is executed: ping 8.8.8.8
```
🎉🎉🎉

## Contributions 🤝

Contributions are welcome! Please feel free to submit a Pull Request or open an issue. 🙏✨

## License 📜

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details. 📄

---

With KSC, managing and executing your frequently used shell commands has never been easier. Enjoy a seamless and interactive experience right from your terminal! 🚀💻🎉

---
