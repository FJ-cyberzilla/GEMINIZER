# Geminizer CLI

Geminizer is an AI-powered automation framework designed to streamline your development and creative workflows. With an interactive CLI, it allows for task orchestration, file analysis, prompt enhancement, and more.

## 🚀 Features

- **Interactive CLI**: Stylized terminal interface using `rich` and `typer`.
- **Task Orchestration**: Easily register and run automated tasks.
- **AI Prompt Enhancement**: Leverage Gemini API to refine your creative prompts.
- **File Analysis**: Quickly analyze project files for metrics.
- **User-Friendly Setup**: Guided onboarding for environment configuration.

## 🛠️ Getting Started

### Prerequisites

- Python 3.14+
- `uv` (recommended for dependency management)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/FJ-cyberzilla/GEMINIZER.git
   cd GEMINIZER
   ```

2. Install dependencies:
   ```bash
   uv sync
   ```

### Configuration

Geminizer requires a Google Gemini API Key to utilize AI-powered features.

1. Obtain a key from [Google AI Studio](https://aistudio.google.com/).
2. Set the environment variable in your terminal session:
   ```bash
   export GOOGLE_API_KEY='your_api_key_here'
   ```
   *For persistent setup, add this line to your shell configuration file (e.g., `~/.bashrc` or `~/.zshrc`).*

## 📖 Usage

Geminizer is controlled via `make` commands:

| Command | Description |
| :--- | :--- |
| `make run` | Runs registered automation tasks. |
| `make chat` | Opens the interactive AI chat interface. |
| `make lint` | Runs type checks (mypy). |
| `make test` | Runs the test suite (pytest). |
| `make clean` | Removes temporary files. |

### Interactive Chat
Run `make chat` to engage with the AI prompt enhancer. Follow the on-screen instructions, and press `Ctrl+C` to exit the chat session.

## 🧪 Development

### Running Tests
To ensure project integrity, run the test suite:
```bash
make test
```

### Type Checking
To verify type safety across the project:
```bash
make lint
```

## 🤝 Contributing

We welcome contributions! Please follow the standard workflow:
1. Fork the repo.
2. Create a feature branch.
3. Commit your changes.
4. Push to the branch and submit a Pull Request.

Ensure all new code is type-checked and tested before submission.
