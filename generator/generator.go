package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	copyFile("thirdparty/cimguizmo/ImGuizmo/ImGuizmo.cpp", "dist/ImGuizmo/ImGuizmo.cpp")
	copyFile("thirdparty/cimguizmo/ImGuizmo/ImGuizmo.h", "dist/ImGuizmo/ImGuizmo.h")

	copyFile("thirdparty/cimguizmo/cimguizmo.cpp", "dist/cimguizmo/cimguizmo.cpp")
	copyFile("thirdparty/cimguizmo/cimguizmo.h", "dist/cimguizmo/cimguizmo.h")

	copyFile("thirdparty/cimgui/cimgui.h", "dist/cimgui/cimgui.h")

	copyFile("thirdparty/imgui/imconfig.h", "dist/imgui/imconfig.h")
	copyFile("thirdparty/imgui/imgui.h", "dist/imgui/imgui.h")
	copyFile("thirdparty/imgui/imgui_internal.h", "dist/imgui/imgui_internal.h")
}

func copyFile(src, dst string) {
	_ = os.MkdirAll(filepath.Dir(dst), 0750)
	srcFile, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer dstFile.Close()

	_, err = srcFile.WriteTo(dstFile)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Copied file %s to %s\n", src, dst)
}
