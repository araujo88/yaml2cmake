cmake_minimum_required(VERSION 3.14)
project(example_test)

# Set the C++ standard
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED True)

# Include FetchContent module for GoogleTest
include(FetchContent)

# Force the use of the shared CRT on Windows
set(gtest_force_shared_crt ON CACHE BOOL "" FORCE)

# Optionally set the MSVC runtime library to DLL
set(CMAKE_MSVC_RUNTIME_LIBRARY "MultiThreaded$<$<CONFIG:Debug>:Debug>DLL")

# Declare GoogleTest as fetch content
FetchContent_Declare(
  googletest
  URL https://github.com/google/googletest/archive/refs/tags/release-1.11.0.zip
)

# Make GoogleTest available for use
FetchContent_MakeAvailable(googletest)

# Create your test executable
add_executable(example_test example_test.cpp)  # Replace example_test.cpp with your actual test file(s)
target_link_libraries(example_test gtest_main)  # Link against gtest_main for GoogleTest's main function

# If your project or tests require threading, link against Threads::Threads
find_package(Threads REQUIRED)
target_link_libraries(example_test Threads::Threads)

# Add the tests to CMake's testing framework
include(GoogleTest)
gtest_discover_tests(example_test)
