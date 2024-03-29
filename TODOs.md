# TODOs

### Feature Enhancements

- [ ] Add support for specifying additional CMake options in the YAML configuration (e.g., compiler flags).
- [ ] Implement functionality to automatically detect and list source files without needing explicit patterns in YAML.
- [ ] Extend external project integration to support other version control systems beyond Git (e.g., SVN).

### Code Quality and Maintenance

- [ ] Refactor the codebase for better modularity and readability.
- [ ] Set up a linter and formatter for consistent code style (e.g., `gofmt` and `golint`).
- [ ] Increase test coverage for existing and new features.
- [ ] Create a Dockerfile to containerize the application, ensuring consistent execution environments.

### Documentation and Examples

- [ ] Update the `README.md` to include more comprehensive usage examples.
- [ ] Add a `TROUBLESHOOTING.md` guide for common issues and fixes.
- [ ] Create detailed documentation for the YAML configuration schema.
- [ ] Provide example projects showing how to use yaml2cmake in various scenarios.

### User Interface and Experience

- [ ] Develop a simple web interface for generating `CMakeLists.txt` files without local setup.
- [ ] Add verbose and quiet mode options to the command-line interface for better user control over output.

### Community Engagement and Growth

- [ ] Establish a contribution guideline that encourages community involvement.
- [ ] Set up a discussion forum or use GitHub Discussions for community support and feature requests.
- [ ] Regularly review and address community feedback and pull requests.

### Automation and Integration

- [ ] Implement Continuous Integration (CI) to automatically run tests on code pushes and pull requests.
- [ ] Automate the release process, including tagging and generating changelogs.
- [ ] Explore plugins or extensions for popular IDEs to integrate yaml2cmake directly into the development workflow.

### Security and Reliability

- [ ] Audit the code for security vulnerabilities, especially in YAML parsing.
- [ ] Implement error handling and validation for edge cases in YAML configurations.

### Future Directions

- [ ] Investigate the possibility of generating other build system configurations (e.g., Bazel, Meson) from YAML.
- [ ] Conduct a user survey to gather feedback and prioritize future development directions.

Focusing on these areas will not only enhance the functionality and usability of yaml2cmake but also engage the community and ensure the project's long-term sustainability and relevance.
