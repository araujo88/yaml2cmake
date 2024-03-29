# yaml2cmake

[![license](https://img.shields.io/badge/license-MIT-green)](https://raw.githubusercontent.com/araujo88/yaml2cmake/main/LICENSE)
[![build](https://github.com/araujo88/yaml2cmake/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/araujo88/yaml2cmake/actions/workflows/build.yml)
[![release](https://img.shields.io/github/v/release/araujo88/yaml2cmake)](https://github.com/araujo88/yaml2cmake/releases)

yaml2cmake simplifies C++ project setups by generating `CMakeLists.txt` files from easy-to-write YAML configurations. It streamlines the process of integrating external dependencies, managing project settings, and ensuring project portability and scalability with minimal effort.

## Features

- **Simple YAML to CMake**: Convert intuitive YAML files into `CMakeLists.txt` configurations.
- **External Dependencies**: Easily include external projects from GitHub repositories.
- **Customizable Setup**: Define project name, version, C++ standards, and more through YAML.
- **Installation Rules**: Automatically generate installation rules for your projects.

## Getting Started

### Prerequisites

- Go 1.22 or higher

### Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/araujo88/yaml2cmake.git
   ```

2. Navigate to the project directory:

   ```bash
   cd yaml2cmake
   ```

3. Build the project (optional):

   ```bash
   go build -o yaml2cmake
   ```

### Usage

1. Create a `config.yaml` file in the project directory with your project configurations. See [Configuration](#configuration) for more details.

2. Run yaml2cmake:

   ```bash
   ./yaml2cmake
   ```

   This will generate a `CMakeLists.txt` file in the same directory.

## Configuration

Your `config.yaml` should look something like this:

```yaml
project_name: example_project
version: 1.0
cpp_standard: 11
library_type: static
include_dirs:
  - include
source_patterns:
  - src/*.cpp
install: true
external_projects:
  - name: example_external_project
    git_repo: https://github.com/user/example_external_project.git
    tag: main
```

For a detailed explanation of each configuration option, refer to [Configuration Options](#configuration-options).

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

Please refer to [CONTRIBUTING.md](CONTRIBUTING.md) for our contribution guidelines.

## License

Distributed under the MIT License. See `LICENSE` for more information.

### Configuration Options

The `config.yaml` file is central to how yaml2cmake functions, allowing users to define the specifics of their C++ project's build configuration with ease. Below, each configuration option available in `config.yaml` is detailed, explaining its purpose, expected values, and format.

#### `project_name`

- **Description**: Specifies the name of your C++ project.
- **Type**: String
- **Example**: `project_name: example_project`

#### `version`

- **Description**: Defines the version of your project. This can be used within CMake to set version-specific properties or requirements.
- **Type**: String
- **Example**: `version: 1.0`

#### `cpp_standard`

- **Description**: Sets the C++ standard version the project should comply with. Common values include `11`, `14`, `17`, `20`, etc.
- **Type**: Integer
- **Example**: `cpp_standard: 17`

#### `library_type`

- **Description**: Determines whether the project should be built as a static or shared library.
- **Type**: String (`static` or `shared`)
- **Example**: `library_type: static`

#### `include_dirs`

- **Description**: A list of directories where CMake should look for header files. These paths are typically relative to the project root.
- **Type**: List of strings
- **Example**:
  ```yaml
  include_dirs:
    - include
    - third_party/include
  ```

#### `source_patterns`

- **Description**: File patterns to identify source files for the project. These patterns are usually relative paths with wildcards.
- **Type**: List of strings
- **Example**:
  ```yaml
  source_patterns:
    - src/*.cpp
    - src/**/*.cpp
  ```

#### `install`

- **Description**: Indicates whether install rules should be generated. Useful for projects intended to be installed on the target system.
- **Type**: Boolean (`true` or `false`)
- **Example**: `install: true`

#### `external_projects`

- **Description**: Defines external projects to be included as dependencies. Each external project can specify a name, git repository URL, and an optional tag (branch, tag, or commit).
- **Type**: List of objects
- **Example**:
  ```yaml
  external_projects:
    - name: external_project
      git_repo: https://github.com/user/external_project.git
      tag: main
    - name: another_dependency
      git_repo: https://github.com/user/another_dependency.git
      tag: v1.2.3
  ```

#### `output_type`

- **Description**: Specifies the type of output the project should generate. It determines whether yaml2cmake configures the project to build an executable binary or a library. For libraries, further specification of the library type (`static` or `shared`) is required.
- **Type**: String (`executable` or `library`)
- **Example**:
  - To generate an executable:
    ```yaml
    output_type: executable
    ```
  - To generate a library (note: `library_type` is also needed for libraries):
    ```yaml
    output_type: library
    library_type: static
    ```

Each option in the `config.yaml` file is designed to be straightforward, making it easy for users to configure their projects without deep knowledge of CMake's syntax or complexities. By adjusting these options, users can quickly scaffold a `CMakeLists.txt` that meets their project's specific needs, from simple applications to complex ones with multiple external dependencies.

## Example Project: Running Your First Build with yaml2cmake

This guide will walk you through the process of using yaml2cmake to configure, build, and run a simple C++ "Hello, World!" project. Follow these steps to see how yaml2cmake simplifies the CMake configuration process.

### Step 1: Generating the CMake Configuration

First, ensure you have your `config.yaml` in the project directory and that it is configured correctly for your project. Then, run yaml2cmake to generate your `CMakeLists.txt` file:

```bash
go run main.go
```

This command reads your `config.yaml` and produces a `CMakeLists.txt` tailored to your project's specifications.

### Step 2: Configuring the Project with CMake

Next, use CMake to configure your project. This step creates the build system in a new directory called `build`:

```bash
cmake -S . -B build
```

### Step 3: Building the Project

With the build system configured, you can now compile your project:

```bash
cmake --build build
```

This command builds the executable defined in your `CMakeLists.txt`, placing it within the `build` directory.

### Step 4: Running the Executable

Finally, run your compiled executable:

```bash
./build/example_project
```

### Expected Output

If everything is set up correctly, you will see something similar to the following output in your terminal:

```
[2024-03-29 17:13:56] [INFO] [/home/leonardo/Github/yaml2cmake/src/main.cpp:6] Hello, world!
```

This output indicates that your "Hello, World!" program has run successfully, demonstrating how yaml2cmake facilitates the process of configuring, building, and running C++ projects with ease.
