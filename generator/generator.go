package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	copyFile("thirdparty/cimguizmo/cimguizmo.h", "dist/include/cimguizmo.h")
	copyFile("thirdparty/cimgui/cimgui.h", "dist/include/cimgui.h")

	copyFile("thirdparty/cimguizmo/ImGuizmo/ImGuizmo.cpp", "dist/imguizmo/ImGuizmo.cpp")
	copyFile("thirdparty/cimguizmo/ImGuizmo/ImGuizmo.h", "dist/imguizmo/ImGuizmo.h")

	copyFile("thirdparty/cimguizmo/cimguizmo.cpp", "dist/cimguizmo/cimguizmo.cpp")
	copyFile("thirdparty/cimguizmo/cimguizmo.h", "dist/cimguizmo/cimguizmo.h")
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
