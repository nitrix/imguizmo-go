package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	copyFile("thirdparty/cimguizmo/cimguizmo.h", "dist/include/cimguizmo.h")
	copyFile("thirdparty/cimgui/cimgui.h", "dist/include/cimgui.h")
}

func copyFile(src, dst string) error {
	_ = os.MkdirAll(filepath.Dir(dst), 0750)
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = srcFile.WriteTo(dstFile)
	if err != nil {
		return err
	}

	fmt.Printf("Copied file %s to %s\n", src, dst)

	return nil
}
