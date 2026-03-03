# TaskCLI

> A fast, minimal, keyboard-driven task manager that lives in your terminal.

Built with [Go](https://go.dev/) and the [Charm](https://charm.sh/) TUI stack — no database, no cloud, no account. Just a JSON file and your terminal.

---

## Preview

```
TaskCLI
Your minimal task manager
══════════════════════════════════════════════════

Progress: ████████████░░░░░░░░░░░░░░░░░░  2/5 (40%)

▌ [ ] Buy groceries
  [ ] Read Clean Code
  [x] Morning run
  [x] Review pull requests
  [ ] Write tests

─────────────────────────────────────────────────
↑/↓: Navigate • N: New • E: Edit • M: Mark • D: Delete • C: Clear • Q: Quit
```

---

## Features

- **Add, edit, delete** tasks instantly from the keyboard
- **Mark complete/incomplete** with a single key
- **Progress bar** showing completed vs total tasks
- **Clear all completed** tasks in one keystroke
- **Persistent storage** — tasks survive restarts, stored in `~/.taskcli.json`
- **Zero config** — works out of the box, no setup needed
- **Cross-platform** — Windows, macOS, Linux

---

## Installation

### Option 1 — Download Pre-built Binary (no Go required)

Head to the [Releases](https://github.com/vaForge/GoTaskCli/releases) page and grab the binary for your OS. No Go installation needed.

| OS      | Architecture             | File                        |
| ------- | ------------------------ | --------------------------- |
| Windows | x86-64                   | `taskcli-windows-amd64.exe` |
| macOS   | Apple Silicon (M1/M2/M3) | `taskcli-macos-arm64`       |
| macOS   | Intel                    | `taskcli-macos-amd64`       |
| Linux   | x86-64                   | `taskcli-linux-amd64`       |
| Linux   | ARM64                    | `taskcli-linux-arm64`       |

**Windows**

1. Download `taskcli-windows-amd64.exe`
2. Rename it to `taskcli.exe` (optional)
3. Double-click or run from any terminal:
   ```powershell
   .\taskcli.exe
   ```
4. To run it from anywhere, move it to a folder on your `PATH`:
   ```powershell
   # Run as administrator
   Move-Item taskcli.exe C:\Windows\System32\taskcli.exe
   ```

**macOS / Linux**

```bash
# 1. Make it executable
chmod +x taskcli-linux-amd64   # or taskcli-macos-arm64 etc.

# 2. Run it directly
./taskcli-linux-amd64

# 3. (Optional) Move to PATH so you can run it from anywhere
sudo mv taskcli-linux-amd64 /usr/local/bin/taskcli
taskcli
```

> macOS users: if Gatekeeper blocks the binary, run:
> `xattr -c taskcli-macos-arm64` then try again.

---

### Option 2 — Install with Go

Requires [Go 1.21+](https://go.dev/dl/) installed.

```bash
go install github.com/vaForge/GoTaskCli@latest
```

Then run:

```bash
taskcli
```

---

### Option 3 — Build from Source

```bash
# 1. Clone the repository
git clone https://github.com/vaForge/GoTaskCli.git
cd taskcli

# 2. Install dependencies
go mod tidy

# 3. Build the binary
go build -o taskcli .

# 4. (Optional) Move to PATH
# Linux/macOS:
sudo mv taskcli /usr/local/bin/

# Windows (PowerShell as admin):
Move-Item taskcli.exe C:\Windows\System32\
```

---

## Keybindings

| Key            | Action                      |
| -------------- | --------------------------- |
| `↑` / `k`      | Move cursor up              |
| `↓` / `j`      | Move cursor down            |
| `N`            | Add a new task              |
| `E`            | Edit the selected task      |
| `M` / `Enter`  | Toggle complete/incomplete  |
| `D`            | Delete the selected task    |
| `C`            | Clear all completed tasks   |
| `Q` / `Ctrl+C` | Quit                        |
| `Esc`          | Cancel input (while typing) |
| `Enter`        | Save input (while typing)   |

---

## Architecture

TaskCLI follows the **Elm Architecture** (Model-Update-View), which is the pattern enforced by the [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework.

```
┌──────────────────────────────────────────────────────┐
│                        main.go                        │
│  Entry point — loads tasks from disk, starts the TUI  │
└────────────────────────┬─────────────────────────────┘
                         │
           ┌─────────────▼──────────────┐
           │        storage/            │
           │  ┌─────────┐ ┌──────────┐  │
           │  │ task.go │ │ file.go  │  │
           │  │  Task   │ │  Load/   │  │
           │  │  struct │ │  Save    │  │
           │  └─────────┘ └──────────┘  │
           └─────────────┬──────────────┘
                         │
           ┌─────────────▼──────────────┐
           │           ui/              │
           │  ┌──────────┐ ┌─────────┐  │
           │  │ model.go │ │styles.go│  │
           │  │  Model   │ │lipgloss │  │
           │  │  Update  │ │ styles  │  │
           │  │  View    │ │         │  │
           │  └──────────┘ └─────────┘  │
           └────────────────────────────┘
```

### Package Breakdown

#### `main.go`

- Loads tasks from disk via `storage.LoadTasks()`
- Creates the initial UI model with `ui.NewModel()`
- Hands control to the Bubble Tea runtime via `tea.NewProgram().Run()`

#### `storage/task.go`

Defines the core data structure:

```go
type Task struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}
```

#### `storage/file.go`

Handles all disk I/O:

- `LoadTasks()` — reads `~/.taskcli.json`, returns an empty list if the file doesn't exist yet
- `SaveTasks()` — marshals the current task slice to pretty-printed JSON and writes it to `~/.taskcli.json`

#### `ui/model.go`

Implements the Bubble Tea **Model-Update-View** loop:

| Function     | Responsibility                                                          |
| ------------ | ----------------------------------------------------------------------- |
| `NewModel()` | Initialises the app state with loaded tasks and a configured text input |
| `Init()`     | Starts the cursor blinking for the text input                           |
| `Update()`   | Handles all keyboard events and mutates state accordingly               |
| `View()`     | Renders the entire terminal UI as a single string                       |

App states (`AppState`):

| State       | Description                                  |
| ----------- | -------------------------------------------- |
| `StateList` | Default — browsing the task list             |
| `StateNew`  | Text input is open for a new task            |
| `StateEdit` | Text input is open to edit the selected task |

#### `ui/styles.go`

All visual styles defined using [Lipgloss](https://github.com/charmbracelet/lipgloss). Centralised here so colours and borders are easy to change without touching logic.

---

## Data Storage

Tasks are stored as a plain JSON file at:

| OS      | Path                           |
| ------- | ------------------------------ |
| Linux   | `~/.taskcli.json`              |
| macOS   | `~/.taskcli.json`              |
| Windows | `C:\Users\<you>\.taskcli.json` |

Example file:

```json
[
  {
    "id": 1,
    "title": "Buy groceries",
    "completed": false
  },
  {
    "id": 2,
    "title": "Morning run",
    "completed": true
  }
]
```

You can back this file up, sync it via Dropbox/iCloud, or version-control it freely.

---

## Dependencies

| Package                                                                 | Purpose                                      |
| ----------------------------------------------------------------------- | -------------------------------------------- |
| [`charmbracelet/bubbletea`](https://github.com/charmbracelet/bubbletea) | TUI framework (Elm Architecture event loop)  |
| [`charmbracelet/bubbles`](https://github.com/charmbracelet/bubbles)     | Reusable TUI components (text input)         |
| [`charmbracelet/lipgloss`](https://github.com/charmbracelet/lipgloss)   | Terminal styling (colours, borders, padding) |

---

## Contributing

Pull requests are welcome.

```bash
# Fork → clone → create a branch
git checkout -b feature/my-improvement

# Make your changes, then run
go build ./...

# Submit a PR
```

---

## Releasing (maintainers)

Binaries for all platforms are built and published automatically by GitHub Actions whenever you push a version tag.

```bash
git tag v1.0.0
git push origin v1.0.0
```

This triggers [`.github/workflows/release.yml`](.github/workflows/release.yml) which cross-compiles for:

- `linux/amd64`, `linux/arm64`
- `darwin/amd64` (Intel Mac), `darwin/arm64` (Apple Silicon)
- `windows/amd64`

and uploads every binary to the GitHub Release automatically.
