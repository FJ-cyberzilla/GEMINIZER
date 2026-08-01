# Geminizer Makefile

# Colors for aesthetic output
GREEN := \033[32m
BLUE := \033[34m
YELLOW := \033[33m
BOLD := \033[1m
RESET := \033[0m

.PHONY: help install run lint test clean

help: ## Show this help message
	@echo "$(BOLD)Geminizer CLI$(RESET)"
	@echo "Usage: make $(BLUE)<command>$(RESET)"
	@echo ""
	@echo "$(BOLD)Commands:$(RESET)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(BLUE)%-15s$(RESET) %s\n", $$1, $$2}'

install: ## Install dependencies (already installed globally)
	@echo "$(GREEN)Dependencies are already installed globally.$(RESET)"

run: ## Run all tasks via the CLI
	@echo "$(GREEN)Running Geminizer...$(RESET)"
	@PYTHONPATH=. python3 src/geminizer_python/cli.py run

lint: ## Run type checking (mypy)
	@echo "$(YELLOW)Running type checks...$(RESET)"
	@python3 -m mypy --strict --explicit-package-bases src/ tests/

test: ## Run tests
	@echo "$(GREEN)Running tests...$(RESET)"
	@PYTHONPATH=. python3 -m pytest tests/

clean: ## Clean up temporary files
	@echo "$(YELLOW)Cleaning up...$(RESET)"
	@find . -type d -name "__pycache__" -exec rm -rf {} +
	@rm -rf .pytest_cache/ .coverage .mypy_cache/
