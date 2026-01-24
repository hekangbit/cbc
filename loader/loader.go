package loader

import "fmt"

type LibraryLoader struct {
	LoadPath []string
}

func LoadLibrary(path string) {
	fmt.Println("LoadLibrary: ", path)
}

func (loader *LibraryLoader) AddLoadPath(path string) {
	loader.LoadPath = append(loader.LoadPath, path)
}
