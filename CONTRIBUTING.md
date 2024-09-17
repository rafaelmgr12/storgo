
# Contributing to Storgo

Thank you for considering contributing to Storgo, a decentralized, content-addressable file storage system! We appreciate your interest in helping us build a robust and reliable system. Please follow the guidelines outlined below to contribute effectively.

## Getting Started

1. **Fork the Repository**:
   - Navigate to the [Storgo repository](https://github.com/rafaelmgr12/storgo) and click on the "Fork" button to create your own copy of the repository.

2. **Clone Your Fork**:
   - Clone your fork locally by running:

     ```bash
     git clone https://github.com/YOUR_USERNAME/storgo.git
     cd storgo
     ```

3. **Set Upstream Remote**:
   - Add the original repository as the upstream remote to keep your fork up-to-date:

     ```bash
     git remote add upstream https://github.com/rafaelmgr12/storgo.git
     ```

4. **Install Dependencies**:
   - Ensure you have Go installed ([Go installation instructions](https://golang.org/doc/install)).
   - Run the following to install dependencies:

     ```bash
     go mod tidy
     ```

5. **Build and Test**:
   - Build the project:

     ```bash
     go build ./...
     ```

   - Run the tests to verify everything is working correctly:

     ```bash
     go test ./...
     ```

## Contribution Guidelines

### 1. Bug Reports and Feature Requests

If you discover a bug or have an idea for a new feature, please [open an issue](https://github.com/rafaelmgr12/storgo/issues) before working on it. This helps avoid duplicate work and gives us a chance to discuss the changes or improvements.

When submitting an issue:

- Clearly describe the problem or feature request.
- Provide steps to reproduce the issue if applicable.
- Suggest a potential solution or describe the feature in detail.

### 2. Pull Requests

When you're ready to submit your contribution, follow these steps:

1. **Create a New Branch**:
   - Always create a new branch for your changes. Use a descriptive name for your branch:

     ```bash
     git checkout -b feature/short-description
     ```

2. **Make Your Changes**:
   - Implement your changes in the new branch. Ensure your code adheres to the project's coding standards.

3. **Test Your Changes**:
   - Run the existing test suite and, if needed, write new tests for the code you've added:

     ```bash
     go test ./...
     ```

4. **Commit Your Changes**:
   - Write meaningful and descriptive commit messages:

     ```bash
     git commit -m "Add short description of the changes"
     ```

5. **Push Your Changes**:
   - Push your branch to your fork:

     ```bash
     git push origin feature/short-description
     ```

6. **Submit a Pull Request**:
   - Go to the original repository and submit a pull request from your branch. Provide a detailed explanation of what your pull request does and reference the issue it addresses (if applicable).

### 3. Code Style

Please adhere to the following coding standards when contributing:

- Write clear, readable Go code.
- Follow the standard Go formatting with `gofmt`.
- Keep functions small and focused. Split large functions into smaller, reusable units when possible.
- Write comments where necessary to clarify the purpose of complex code.
  
### 4. Testing

- Ensure your contribution has adequate test coverage.
- Write unit tests for new code or features.
- Make sure tests pass by running `go test ./...`.

## Additional Notes

- **Be Respectful**: Storgo is a community-driven project, and we value constructive feedback. Be respectful in your communication and code reviews.
  
- **Stay Up-to-Date**: Before starting work on a new feature or bug fix, ensure your branch is up-to-date with the latest changes from the `master` branch:

  ```bash
  git fetch upstream
  git checkout master
  git merge upstream/master
  ```

- **License**: By contributing to Storgo, you agree that your contributions will be licensed under the project's [MIT License](LICENSE).

---

Thank you again for your interest in contributing to Storgo. We look forward to your contributions!
