# 🐳 dockgen

`dockgen` is a Go-based toolchain that parses shell commands, builds intermediate representations (IR), constructs control flow graphs (CFG), and optionally generates Dockerfiles. Inspired by compiler design and Unix tooling philosophy, `dockgen` helps developers analyze, refactor, and containerize CLI workflows.

## ✨ Features

- Parse bash commands into structured IR (supports pipes, redirects, env, options)
- Build Control Flow Graphs (CFG) from parsed commands
- Graph analysis and optimizations (e.g., family grouping, deduplication)
- Output graph structure in JSON or DOT format
- Generate valid Dockerfiles from analyzed command sequences
- Interactive watch mode for debugging commands in real-time

## 📦 Installation

```bash
go install github.com/your-org/dockgen@latest

```
🚀 Usage
```bash
dockgen build < input.json            # Build IR from JSON input
dockgen graph --format=dot            # Output control flow graph
dockgen dockerfile --output=Dockerfile
Available Commands
Command	Description
build	Parse JSON and generate IR
graph	Build and display the control flow graph
dockerfile	Generate Dockerfile from IR graph
--watch	Watch command input and stream live IR output
```

🧠 Architecture
IR Layer: Represents parsed commands as a structured tree (program, args, env, redirects, etc.)

CFG Layer: Control flow graph builder and analysis layer

Analyzer: Performs optimizations like family merging and deduplication

Generators: Code emitters like Dockerfile generator or graph exporters

📝 Example
Input JSON (input.json):

```json
[
  {"cmd": "apt update", "stdout": "", "stderr": ""},
  {"cmd": "apt install curl", "stdout": "", "stderr": ""},
  {"cmd": "apt install wget", "stdout": "", "stderr": ""}
]
```
Run:

```bash
dockgen build < input.json | dockgen graph | dockgen dockerfile
🔧 Development
Written in Go

Uses ANTLR for Bash grammar parsing

CLI built with Cobra

Graph optimizations follow a compiler-inspired pass model
```
📂 Project Structure
```bash
cmd/           # CLI commands
parser/        # Bash grammar and ANTLR integration
ir/            # Intermediate Representation structures
graph/         # CFG and optimizations
gen/           # Dockerfile and output generation
internal/      # Utilities and internal helpers
```

🛠 Future Work
DSL support for retry and postprocessing rules

TTY-aware session replay

Rich visual graph UI

Kubernetes target generation

