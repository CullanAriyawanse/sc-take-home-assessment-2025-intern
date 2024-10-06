package folder

import (
	"errors"
	"slices"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	folderData := make([]Folder, len(f.folders))
	copy(folderData, f.folders)

	var err error

	// Check if source and destination folders are the same
	if name == dst {
		err = errors.New("Error: Cannot move a folder to a child of itself")
		return folderData, err
	}

	var parentFolderPath string
	var foundSrcFolder bool

	// Get path to parent folder
	for _, folder := range folderData {
		if folder.Name == name {
			foundSrcFolder = true
		}
		if folder.Name == dst {
			parentFolderPath = folder.Paths
		}
	}

	// Check if source folder exists
	if !foundSrcFolder {
		err = errors.New("Error: Source folder does not exist")
		return folderData, err
	}

	// Check if destination folder exists
	if parentFolderPath == "" {
		err = errors.New("Error: Destination folder does not exist")
		return folderData, err
	}

	var newFullPath string

	// Change path of subtree
	for i, folder := range folderData {
		if folder.Name == name {
			newFullPath = parentFolderPath + "." + name
			folderData[i].Paths = newFullPath
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
