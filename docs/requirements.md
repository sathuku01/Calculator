## ** 🧮 Go Calculator with UI **

Version: 1.0.0
Author: Thuku Samuel / Clare
Last Updated: 02/02/2026

## **📌 Project Overview**
Description

This project is a full-featured calculator implemented in the Go programming language with a graphical user interface (UI). The calculator will support basic arithmetic operations, a responsive UI, error handling, a history panel, and extensibility for future enhancements.

Goal

This project will enable us to learn industry best practices by letting us focus on clean code structure, clear requirements, and readable naming without excessive complexity. It provides hands-on experience with version control, testing, and error handling in a realistic but manageable context. Because the project is small, it can be safely refactored and code documented, practicing the full development lifecycle used in professional software development.

## ** 📚 Table of Contents **

Goals & Objectives

User Stories

Functional Requirements

Non-Functional Requirements

User Interface (UI) Requirements

Architecture & Technologies

Data Handling & Storage

Error Handling

Testing Requirements

Project Milestones

## 1. 🎯 Goals & Objectives
1.1 Primary Goal

Build a reliable, easy-to-use calculator application written in Go with a modern UI.

1.2 Secondary Goals

Ensure accurate mathematical computations

Provide a clean, intuitive UI

Support both keyboard and mouse interactions

Maintain extensible code architecture

## 2. 👤 User Stories
ID	User Story
US-01	As a user, I can perform basic arithmetic: addition, subtraction, multiplication, division.
US-02	As a user, I can use advanced functions like exponentiation and percentages.
US-03	As a user, I can see a history of my calculations.
US-04	As a user, I can clear my input or reset the calculator.
US-05	As a user, I can navigate using keyboard shortcuts.
US-06	As a developer, I can run unit tests automatically.
US-07	As a developer, I can easily extend the application with new features.

## 3. ⚙️ Functional Requirements
3.1 Calculator Logic

FR-1: Accept numeric input (integer and decimal).

FR-2: Support basic operations: +, −, ×, ÷.

FR-3: Support advanced operations: √, ^ (power), % (percentage), ± (negate).

FR-4: Follow correct operator precedence.

FR-5: Allow chaining of operations.

FR-6: Display results immediately upon request.

FR-7: Maintain a calculable state internally.

3.2 UI Interaction

FR-8: Buttons for all supported functions.

FR-9: Keyboard input support for digits, operators, and function keys.

FR-10: History display of previous calculations.

## 3.3 Data Management

FR-11: Store history in memory.

FR-12: Optionally export history to a file (e.g., .txt or .csv).

## 4. 🧩 Non-Functional Requirements
4.1 Performance

NFR-1: UI must update within 100ms of interaction.

NFR-2: Calculation results must be computed within 50ms.

## 4.2 Usability

NFR-3: UI must be intuitive and responsive.

NFR-4: Keyboard shortcuts should mirror standard calculator conventions.

## 4.3 Reliability

NFR-5: Calculator must handle invalid input gracefully.

NFR-6: No crashes due to malformed input.

## 4.4 Portability

NFR-7: Application should run on Windows, macOS, and Linux.

## 5. 🎨 User Interface (UI) Requirements
## 5.1 Layout

UI-1: A display screen showing input and results.

UI-2: Button grid for numbers 0–9.

UI-3: Buttons for operations and functions.

UI-4: History section (scrollable) showing recent calculations.

## 5.2 Visual Style

UI-5: Use a clean, modern color palette.

UI-6: Clearly distinguish numeric buttons from operators.

## 5.3 Accessibility

UI-7: All buttons must have keyboard focus states.

UI-8: Provide alternative text for icons and non-text elements.

## 5.4 Responsiveness

UI-10: Resize gracefully on different screen sizes.

## 6. 🏗 Architecture & Technologies
## 6.1 Architecture

Layered Architecture:

UI Layer (handles visuals and events)

Logic Layer (performs calculation)

Data Layer (history, persistence)

## 6.2 Technologies

Language: Go (>= 1.20)

UI Library: Either

Fyne

Go Qt/Bindi

WebView with HTML/CSS/JS

Testing: Go’s testing package

Build Tools: Go Modules

## 7. 💾 Data Handling & Storage
## 7.1 In-Memory Storage

History buffer storing last N operations.

## 8. 🚨 Error Handling

EH-1: Handle division by zero gracefully.

EH-2: Prompt user on invalid expressions.

EH-3: UI must display error messages clearly.

EH-4: Prevent incorrect states or crashes.

## 9. 🧪 Testing Requirements
## 9.1 Unit Tests

Test each operation.

Test operator precedence.

Test error handling.

## 9.2 UI Tests

Automated UI interaction tests.

## 9.3 Acceptance Tests

Validate user stories.