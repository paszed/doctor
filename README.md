# 🩺 Doctor – Developer Environment Diagnostics

Doctor is a CLI tool for automatically checking, diagnosing, and fixing your developer environment.  
It ensures essential tools, databases, language runtimes, and DevOps utilities are installed, configured, and ready to use.

---

## ⚡ Features

- **Interactive mode**: Run `./doctor` to see your environment status in a structured table.
- **Automatic fixes**: Instantly fix missing or misconfigured tools when supported.
- **Check individual tools**: Focus on a single tool with `./doctor check <tool>`.
- **Full diagnostics**: Run `./doctor diagnose` for a comprehensive environment report.
- **Cross-platform support**: Works on macOS, Linux, and Windows (with Homebrew, apt, or equivalent).
- **Dev toolkit coverage**: Includes databases, language runtimes, package managers, CI/CD tools, Kubernetes, and more.

---

## 📦 Installation

Clone the repository:

```bash
git clone https://github.com/paszed/doctor.git
cd doctor
```

### Option 1: Build from source (Go required)

```bash
# macOS / Linux / Windows (if Go is installed)
go build -o doctor ./cmd/doctor
```

This produces a `doctor` binary in the current folder. 
Optionally, add it to your PATH for global usage:

```bash
export PATH=$PATH:$(pwd)
```

### Option 2: Use prebuilt binaries (no Go required)

```bash
# macOS Apple Silicon
./dist/doctor-mac

# macOS Intel
./dist/doctor-mac-intel

# Linux x86_64
./dist/doctor-linux

# Linux ARM64
./dist/doctor-linux-arm

# Windows x86_64
dist\doctor.exe
```

```bash
# Make Unix binaries executable
chmod +x ./dist/doctor-mac ./dist/doctor-mac-intel ./dist/doctor-linux ./dist/doctor-linux-arm
```
---

## 🚀 Usage

### Interactive check

```bash
./doctor
```

### Full environment diagnostics

```bash
./doctor diagnose
```

### Check a single tool

```bash
./doctor check <tool>
```

### Fix all issues or a specific tool

```bash
./doctor fix
./doctor fix <tool>
```

### Check a specific port

```bash
./doctor port 3000
```

### List all available checks

```bash
./doctor list
```

### Show Doctor version

```bash
./doctor version
```

---

## 🧪 Example: Python package managers

```bash
./doctor check pip
./doctor check pipenv
./doctor check poetry
./doctor fix pip
./doctor fix pipenv
./doctor fix poetry
```

---

## 🔧 Supported Tools

**Languages & Runtimes**

- Go
- Java (OpenJDK)
- Node.js / npm / pnpm / yarn / nvm
- Python 3 / pip / pipenv / poetry / pyenv
- Ruby via rbenv

**Databases**

- PostgreSQL
- MySQL
- MongoDB
- SQLite CLI
- Redis CLI
- Oracle SQL CLI (`sqlplus`)

**DevOps / CI/CD / Cloud**

- Docker / Docker Compose
- Kubernetes (`kubectl`), Kind, Minikube
- Helm
- Terraform
- AWS CLI
- Azure CLI
- Google Cloud CLI
- GitHub CLI (`gh`)

**Other Tools**

- Maven
- Ansible

---

## 📝 Notes

- Some fixes (Terraform, Kubernetes tools, Oracle SQL) may require manual installation or elevated permissions.
- For Python packages, `pipx` is recommended for system-managed Python installations.
- All interactive fixes show suggested installation commands when auto-fix isn’t available.

---

## 🤝 Contribution

1. Fork the repository  
2. Add new checks/fixes in `internal/checks` and `internal/fix`  
3. Update `internal/cli/help.go` and the README  
4. Submit a pull request  
