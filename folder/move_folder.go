package folder

import (
	"fmt"
	"slices"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	folderData := make([]Folder, len(f.folders))
	copy(folderData, f.folders)
	// folderData := f.folders

	// var err error

	var parentFolderPath string

	// Get path to parent folder
	for _, folder := range folderData {
		fmt.Printf("Folder name is %s\n", folder.Name)
		if folder.Name == dst {
			parentFolderPath = folder.Paths
			break
		}
	}

	fmt.Printf("Parent folder is %s!\n", parentFolderPath)

	var newFullPath string
	// Change path of subtree
	for i, folder := range folderData {
		if folder.Name == name {
			newFullPath = parentFolderPath + "." + name
			fmt.Printf("Old path is %s\n", folder.Paths)
			folderData[i].Paths = newFullPath
			fmt.Printf("New path is %s\n", folder.Paths)
			break
		}
	}

	// Change path of all children nodes
	for i, folder := range folderData {
		splitString := strings.Split(folder.Paths, ".")

		containsParent := slices.Contains(splitString, name)

		if containsParent && folder.Name != name && folder.Name != dst {
			folderData[i].Paths = newFullPath + "." + folder.Name
			break
		}
	}

	return folderData, nil
}
