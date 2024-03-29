cmake_minimum_required(VERSION 3.10)
project(example_project VERSION 1.0)
set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED True)

# Include the FetchContent module
include(FetchContent)

# Declare the dependency using FetchContent
FetchContent_Declare(
  logtard
  GIT_REPOSITORY https://github.com/araujo88/logtard.git
  GIT_TAG        main
)
# Make the content available
FetchContent_MakeAvailable(logtard)

file(GLOB SOURCES "src/*.cpp" )

# Generate an executable from specified sources
add_executable(example_project ${SOURCES})

target_include_directories(example_project PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/include )

# Optional: Install rules
install(TARGETS example_project
  ARCHIVE DESTINATION lib
  LIBRARY DESTINATION lib
  RUNTIME DESTINATION bin # For Windows DLLs
)
install(DIRECTORY include/ DESTINATION include)

# Link the external library targets to your project
target_link_libraries(example_project PRIVATE logtard)
