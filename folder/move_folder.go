package folder

import (
	"errors"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	folderData := make([]Folder, len(f.folders))
	copy(folderData, f.folders)

	var err error

	var srcFolder Folder
	var dstFolder Folder
	var foundSrcFolder bool
	var foundDstFolder bool

	for _, folder := range folderData {
		if folder.Name == name {
			srcFolder = folder
			foundSrcFolder = true
		}
		if folder.Name == dst {
			dstFolder = folder
			foundDstFolder = true
		}
	}

	// Check if source folder exists
	if !foundSrcFolder {
		err = errors.New("Error: Source folder does not exist")
		return folderData, err
	}

	// Check if destination folder exists
	if !foundDstFolder {
		err = errors.New("Error: Destination folder does not exist")
		return folderData, err
	}

	// Check if source and destination are the same
	if name == dst {
		err = errors.New("Error: Cannot move a folder to itself")
		return folderData, err
	}

	// Check if destination is a child of source
	if strings.HasPrefix(dstFolder.Paths, srcFolder.Paths) {
		err = errors.New("Error: Cannot move a folder to a child of itself")
		return folderData, err
	}

	// Check if the source and destination belong to the same organization
	if srcFolder.OrgId != dstFolder.OrgId {
		err = errors.New("Error: Cannot move a folder to a different organization")
		return folderData, err
	}

	var newFullPath string
	parentFolderPath := dstFolder.Paths

	// Change path of the source folder
	for i, folder := range folderData {
		if folder.Name == name {
			newFullPath = parentFolderPath + "." + name
			folderData[i].Paths = newFullPath
			break
		}
	}

	// Change path of all children folders
	for i, folder := range folderData {
		if strings.HasPrefix(folder.Paths, srcFolder.Paths) {
			newSubPath := strings.Replace(folder.Paths, srcFolder.Paths, newFullPath, 1)
			folderData[i].Paths = newSubPath
		}
	}

	return folderData, nil
}
