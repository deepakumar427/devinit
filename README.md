# 🚀 devinit

> A lightning-fast, interactive CLI tool to scaffold production-ready projects in seconds.

Tired of writing the same `.gitignore`, `Dockerfile`, and GitHub Actions workflow every time you start a new project? `devinit` automates the boring parts of project setup — giving you an interactive terminal UI to pick your language and features, then handling the rest from Git init to pushing a live GitHub repository.

---

## ✨ Features

- **Interactive UI** — Clean, arrow-key navigation to build your project setup
- **Language Aware** — Automatically generates the correct `.gitignore`, Dockerfiles, and CI/CD pipelines for Go, Node.js, and Python
- **Instant Git & GitHub** — Runs `git init`, stages your initial commit, creates a remote repo via the GitHub CLI, and pushes automatically
- **Modular Architecture** — Built with a clean `Step` interface, making it easy to add new scaffolding steps in the future

---

## 🛠 Prerequisites

| Tool | Version | Required |
|------|---------|----------|
| [Go](https://go.dev/) | 1.20+ | ✅ Yes |
| [Git](https://git-scm.com/) | Any | ✅ Yes |
| [GitHub CLI (`gh`)](https://cli.github.com/) | Any (authenticated) | ✅ Yes |
| [Docker](https://www.docker.com/) | Any | ⚡ Optional |

> **Not sure if you're set up?** Run the built-in doctor command:
> ```bash
> devinit doctor
> ```

---

## 📦 Installation

Clone the repository and install globally via Go:

```bash
git clone https://github.com/yourusername/devinit.git
cd devinit
go install ./...
```

> **Note:** Make sure `~/go/bin` is in your system's `PATH`.

---

## 💻 Usage

Start scaffolding a new project with:

```bash
devinit init <your-project-name>
```

The CLI will walk you through a few quick prompts:

1. **Which programming language are you using?**
2. **Do you want a Dockerfile?**
3. **Do you want to set up GitHub Actions CI/CD?**
4. **Do you want to automatically create and link a GitHub repository?**

Hit enter and watch the green checkmarks roll in. ✅

---

## 🏗 Architecture

`devinit` is built with a focus on clean design and extensibility.

| Component | Role |
|-----------|------|
| [Cobra](https://github.com/spf13/cobra) | Command routing and CLI structure |
| [Survey](https://github.com/AlecAivazis/survey) | Interactive terminal prompts |
| `Step` Interface | Every scaffolding action (Git, Docker, CI/CD) implements `Name()` and `Run()` |

Adding a new feature — like a Terraform step or a Kubernetes manifest generator — is as simple as creating a new struct that satisfies the `Step` interface.

---

## 🤝 Contributing

Contributions are welcome! Want to add support for Rust, Java, or React? Here's how:

1. Fork the project
2. Create your feature branch
```bash
   git checkout -b feature/AmazingFeature
```
3. Commit your changes
```bash
   git commit -m 'Add some AmazingFeature'
```
4. Push to the branch
```bash
   git push origin feature/AmazingFeature
```
5. Open a Pull Request

---

## 📝 License

Distributed under the **MIT License**. See [`LICENSE`](./LICENSE) for more information.
