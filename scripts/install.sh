#!/bin/bash

# Geminizer Enterprise Installation Script
set -e

echo "🚀 Installing Geminizer Enterprise..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

# Detect platform
detect_platform() {
    case "$(uname -s)" in
        Linux*)     platform=linux;;
        Darwin*)    platform=macos;;
        CYGWIN*)    platform=windows;;
        MINGW*)     platform=windows;;
        *)          platform=unknown;;
    esac
    echo $platform
}

# Check dependencies
check_dependencies() {
    echo "📋 Checking dependencies..."
    
    # Check Go
    if command -v go &> /dev/null; then
        echo -e "${GREEN}✅ Go installed${NC}"
    else
        echo -e "${RED}❌ Go not installed${NC}"
        echo "Please install Go 1.21+ from https://golang.org/dl/"
        exit 1
    fi
    
    # Check Node.js
    if command -v node &> /dev/null; then
        echo -e "${GREEN}✅ Node.js installed${NC}"
    else
        echo -e "${RED}❌ Node.js not installed${NC}"
        echo "Please install Node.js 18+ from https://nodejs.org/"
        exit 1
    fi
    
    # Check Docker (optional)
    if command -v docker &> /dev/null; then
        echo -e "${GREEN}✅ Docker installed${NC}"
    else
        echo -e "${YELLOW}⚠️ Docker not installed (optional for containerized deployment)${NC}"
    fi
}

# Install backend
install_backend() {
    echo "🔧 Installing backend dependencies..."
    cd backend
    go mod download
    go mod verify
    cd ..
    echo -e "${GREEN}✅ Backend dependencies installed${NC}"
}

# Install frontend
install_frontend() {
    echo "🎨 Installing frontend dependencies..."
    cd frontend
    npm ci --silent
    cd ..
    echo -e "${GREEN}✅ Frontend dependencies installed${NC}"
}

# Setup environment
setup_environment() {
    echo "⚙️ Setting up environment..."
    
    if [ ! -f .env ]; then
        cp .env.example .env
        echo -e "${YELLOW}⚠️ Please edit .env file with your configuration${NC}"
    else
        echo -e "${GREEN}✅ Environment file already exists${NC}"
    fi
}

# Security setup
setup_security() {
    echo "🔒 Setting up security..."
    
    # Generate JWT secret if not exists
    if ! grep -q "JWT_SECRET" .env 2>/dev/null; then
        jwt_secret=$(openssl rand -base64 64 | tr -d '\n')
        echo "JWT_SECRET=$jwt_secret" >> .env
        echo -e "${GREEN}✅ JWT secret generated${NC}"
    fi
    
    # Set proper permissions
    chmod 600 .env 2>/dev/null || true
    echo -e "${GREEN}✅ Security setup complete${NC}"
}

# Platform-specific setup
platform_setup() {
    local platform=$(detect_platform)
    
    case $platform in
        linux)
            echo "🐧 Linux setup..."
            # Additional Linux dependencies
            if command -v apt &> /dev/null; then
                sudo apt update && sudo apt install -y build-essential
            elif command -v yum &> /dev/null; then
                sudo yum groupinstall -y "Development Tools"
            fi
            ;;
        macos)
            echo "🍎 macOS setup..."
            # Additional macOS setup
            if command -v brew &> /dev/null; then
                brew install pkg-config
            fi
            ;;
        windows)
            echo "🪟 Windows setup..."
            # Additional Windows setup
            if command -v choco &> /dev/null; then
                choco install make
            fi
            ;;
    esac
}

# Main installation
main() {
    echo "🎯 Starting Geminizer Enterprise installation..."
    
    check_dependencies
    platform_setup
    install_backend
    install_frontend
    setup_environment
    setup_security
    
    echo ""
    echo -e "${GREEN}🎉 Installation complete!${NC}"
    echo ""
    echo "Next steps:"
    echo "1. Edit .env file with your configuration"
    echo "2. Run 'make dev' to start development server"
    echo "3. Access the application at http://localhost:3000"
    echo ""
    echo "For production deployment, run 'make deploy'"
}

# Run installation
main "$@"
