# 🚀 devinit 

> A lightning-fast, interactive CLI tool to scaffold production-ready projects in seconds.

Ever get tired of writing the same `.gitignore`, `Dockerfile`, and GitHub Actions workflow every time you start a new project? Yeah, me too. 

I built `devinit` to automate the boring parts of project setup. Instead of copying and pasting boilerplate for the 100th time, `devinit` gives you an interactive terminal UI to pick your language and features, and handles the rest—from initializing Git to creating and pushing a remote GitHub repository.

## ✨ Features

- **Interactive UI:** Clean, arrow-key navigation to build your project setup.
- **Language Aware:** Automatically generates the correct `.gitignore`, Dockerfiles, and CI/CD pipelines for **Go**, **Node.js**, and **Python**.
- **Instant Git & GitHub:** Runs `git init`, stages your initial commit, creates a remote repo via the GitHub CLI, and pushes your code automatically.
- **Modular Architecture:** Built in Go using the Cobra CLI framework and a clean `Step` interface, making it incredibly easy to add new scaffolding steps in the future.

## 🛠 Prerequisites

To get the absolute most out of `devinit`, make sure you have the following installed:

- [Go](https://go.dev/) (1.20+)
- [Git](https://git-scm.com/)
- [GitHub CLI (`gh`)](https://cli.github.com/) - Authenticated and ready to go.
- [Docker](https://www.docker.com/) (Optional, but recommended)

Not sure if you have everything? Run the built-in doctor command!
```bash
devinit doctor

📦 Installation
Since devinit is built in Go, you can easily install it globally on your machine.

Clone this repository and run:
git clone [https://github.com/yourusername/devinit.git](https://github.com/yourusername/devinit.git)
cd devinit
go install ./...

(Note: Ensure your ~/go/bin directory is in your system's PATH!)

💻 Usage
Whenever you are ready to start a new project, just open your terminal and run:

devinit init <your-project-name>


The CLI will ask you a few quick questions:

Which programming language are you using?

Do you want a Dockerfile?

Do you want to set up GitHub Actions CI/CD?

Do you want to automatically create and link a GitHub repository?

Hit enter, and watch the green checkmarks roll in.

🏗 Under the Hood
This project was built with a focus on clean system design and extensibility:

Cobra: Powers the command routing and CLI structure.

Survey: Handles the interactive terminal prompts.

Interface-Driven Design: Every scaffolding action (Git, Docker, CI/CD) implements a Step interface. Adding a new feature (like a Terraform step or a Kubernetes manifest step) is as simple as creating a new struct that satisfies Name() and Run().

🤝 Contributing
Want to add support for Rust, Java, or React? Pull requests are absolutely welcome!

Fork the project.

Create your feature branch (git checkout -b feature/AmazingFeature).

Commit your changes (git commit -m 'Add some AmazingFeature').

Push to the branch (git push origin feature/AmazingFeature).

Open a Pull Request.

📝 License
Distributed under the MIT License. See LICENSE for more information.

### Next Steps:
If you want to add this to your GitHub repository right now, run these commands in your terminal (make sure you are in the `devinit` folder):

```bash
# Open VS Code to create and save the README
code README.md

(Paste the text above into the file and save it, making sure to replace yourusername in the installation link with your actual GitHub username!)

Then, push it up:
git add README.md
git commit -m "docs: add polished human-written README"
git push origin main