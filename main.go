package main

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config represents the structure of the YAML configuration file
type Config struct {
	ProjectName      string            `yaml:"project_name"`
	Version          string            `yaml:"version"`
	CPPStandard      int               `yaml:"cpp_standard"`
	LibraryType      string            `yaml:"library_type"` // static or shared
	IncludeDirs      []string          `yaml:"include_dirs"`
	SourcePatterns   []string          `yaml:"source_patterns"`
	Install          bool              `yaml:"install"`
	ExternalProjects []ExternalProject `yaml:"external_projects"`
	OutputType       string            `yaml:"output_type"`
	Testing          *Testing          `yaml:"testing"`
}

type Testing struct {
	EnableTesting bool   `yaml:"enable_testing"`
	Dir           string `yaml:"dir"`
}

// ExternalProject represents an external project to be added to the CMakeLists
type ExternalProject struct {
	Name    string `yaml:"name"`
	GitRepo string `yaml:"git_repo"`
	Tag     string `yaml:"tag,omitempty"`
}

func main() {
	// Read YAML configuration file
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	// Parse YAML
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	// Generate CMakeLists.txt content
	cmakeContent := generateCMakeLists(config)

	// Write to CMakeLists.txt
	err = os.WriteFile("CMakeLists.txt", []byte(cmakeContent), 0644)
	if err != nil {
		panic(err)
	}
}

func generateCMakeLists(config Config) string {
	var sb strings.Builder

	// Basic project setup
	sb.WriteString("cmake_minimum_required(VERSION 3.10)\n")
	sb.WriteString("project(" + config.ProjectName + " VERSION " + config.Version + ")\n")
	sb.WriteString("set(CMAKE_CXX_STANDARD " + fmt.Sprintf("%d", config.CPPStandard) + ")\n")
	sb.WriteString("set(CMAKE_CXX_STANDARD_REQUIRED True)\n")

	// Include the FetchContent module if there are external projects
	if len(config.ExternalProjects) > 0 {
		sb.WriteString("\n# Include the FetchContent module\ninclude(FetchContent)\n")
		for _, project := range config.ExternalProjects {
			sb.WriteString("\n# Declare the dependency using FetchContent\nFetchContent_Declare(\n")
			sb.WriteString("  " + project.Name + "\n")
			sb.WriteString("  GIT_REPOSITORY " + project.GitRepo + "\n")
			sb.WriteString("  GIT_TAG        " + project.Tag + "\n")
			sb.WriteString(")\n")
			sb.WriteString("# Make the content available\nFetchContent_MakeAvailable(" + project.Name + ")\n")
		}
	}

	// Source files
	sb.WriteString("\nfile(GLOB SOURCES ")
	for _, pattern := range config.SourcePatterns {
		sb.WriteString(`"` + pattern + `" `)
	}
	sb.WriteString(")\n")

	// Determine output type
	if config.OutputType == "executable" {
		sb.WriteString("\n# Generate an executable from specified sources\n")
		sb.WriteString("add_executable(" + config.ProjectName + " ${SOURCES})\n")
	} else if config.OutputType == "library" {
		libType := "STATIC"
		if config.LibraryType == "shared" {
			libType = "SHARED"
		}
		sb.WriteString("\n# Library type\nadd_library(" + config.ProjectName + " " + libType + " ${SOURCES})\n")
	}
	// Include directories
	if len(config.IncludeDirs) > 0 {
		sb.WriteString("\ntarget_include_directories(" + config.ProjectName + " PUBLIC ")
		for _, dir := range config.IncludeDirs {
			sb.WriteString("${CMAKE_CURRENT_SOURCE_DIR}/" + dir + " ")
		}
		sb.WriteString(")\n")
	}

	// Optional: Install rules
	if config.Install {
		sb.WriteString("\n# Optional: Install rules\ninstall(TARGETS " + config.ProjectName + "\n")
		sb.WriteString("  ARCHIVE DESTINATION lib\n")
		sb.WriteString("  LIBRARY DESTINATION lib\n")
		sb.WriteString("  RUNTIME DESTINATION bin # For Windows DLLs\n")
		sb.WriteString(")\ninstall(DIRECTORY include/ DESTINATION include)\n")
	}

	// Assuming the executable or library needs to link against the external projects
	if len(config.ExternalProjects) > 0 {
		sb.WriteString("\n# Link the external library targets to your project\n")
		for _, project := range config.ExternalProjects {
			sb.WriteString("target_link_libraries(" + config.ProjectName + " PRIVATE " + project.Name + ")\n")
		}
	}

	if config.Testing != nil {
		if config.Testing.EnableTesting {
			sb.WriteString("\n# Enable testing with CMake")
			sb.WriteString("\nenable_testing()\n")
		}
		if config.Testing.Dir != "" {
			sb.WriteString("\n# Include the directory where the test code is located")
			sb.WriteString("\nadd_subdirectory(" + config.Testing.Dir + ")\n")
		}
	}

	return sb.String()
}
