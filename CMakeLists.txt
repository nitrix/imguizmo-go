CMAKE_MINIMUM_REQUIRED(VERSION 3.16)
PROJECT(cimguizmo)

SET(CMAKE_CXX_STANDARD 17)

ADD_SUBDIRECTORY(thirdparty/cimgui)

ADD_LIBRARY(imguizmo thirdparty/cimguizmo/ImGuizmo/ImGuizmo.cpp)
TARGET_INCLUDE_DIRECTORIES(imguizmo PRIVATE thirdparty/cimgui/imgui)

ADD_LIBRARY(cimguizmo thirdparty/cimguizmo/cimguizmo.cpp)
TARGET_INCLUDE_DIRECTORIES(cimguizmo PRIVATE thirdparty/cimgui/imgui)
TARGET_INCLUDE_DIRECTORIES(cimguizmo PUBLIC thirdparty/cimgui)
TARGET_INCLUDE_DIRECTORIES(cimguizmo PUBLIC thirdparty/cimguizmo)
TARGET_COMPILE_DEFINITIONS(cimguizmo INTERFACE CIMGUI_DEFINE_ENUMS_AND_STRUCTS)
TARGET_LINK_LIBRARIES(cimguizmo PRIVATE stdc++ imguizmo)

IF(WIN32)
    TARGET_LINK_LIBRARIES(cimguizmo PRIVATE imm32)
ENDIF()

IF(APPLE)
    TARGET_LINK_LIBRARIES(cimguizmo PRIVATE "-framework CoreFoundation")
ENDIF()