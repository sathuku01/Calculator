A feature-rich calculator application written in Go that demonstrates professional software architecture and best practices. This project provides two distinct interfaces—a command-line calculator and a graphical user interface—powered by a shared calculation engine.

# Key Features:

  # Dual Interface Options:
        CLI mode for quick terminal-based calculations
        Modern GUI with Fyne framework featuring a colorful, intuitive calculator layout

  # Robust Calculation Engine:
        Supports basic arithmetic operations: addition, subtraction, multiplication, and division
        Implements proper operator precedence (multiplication and division before addition and subtraction)
        Comprehensive error handling with custom error types

  # Well-Structured Architecture:
        internal/core: Core calculation functions with error handling
        internal/service: Expression tokenization for parsing user input
        internal/parser: Two-pass expression evaluation respecting operator precedence
        internal/ui: GUI components built with Fyne for a polished user experience
        cmd: Separate entry points for CLI and GUI applications

  # Quality Assurance:
        Unit tests for core functionality and tokenization
        Proper input validation and error messages
        Clean, maintainable code structure following Go conventions

# Technology Stack:

    Go 1.22.2
    Fyne v2 for cross-platform GUI

# Perfect for:

    Learning clean code architecture and separation of concerns
    Understanding expression parsing and operator precedence
    Implementing multi-interface applications in Go
    Exploring GUI development with modern frameworks
