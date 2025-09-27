cmake_minimum_required(VERSION 3.10)
project(pyobject_dict)

# Find Python development libraries
find_package(Python3 COMPONENTS Interpreter Development REQUIRED)

# Create executable
add_executable(pyobject_dict main.c)

# Link Python libraries
target_link_libraries(pyobject_dict ${Python3_LIBRARIES})

# Include Python headers
target_include_directories(pyobject_dict PRIVATE ${Python3_INCLUDE_DIRS})

# Set C standard
set_property(TARGET pyobject_dict PROPERTY C_STANDARD 99)
