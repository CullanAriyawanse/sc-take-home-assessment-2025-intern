package folder

import (
	"errors"
	"slices"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	// folderData := GetSampleData()
	folderData := f.folders

	var err error

	childFolders := []Folder{}
	foundFolderDifferentOrg := false
	foundFolderWithOrg := false

	// Split path into slice, if index of parent folder is less than child folder, add child folder to slice
	for _, node := range folderData {
		splitString := strings.Split(node.Paths, ".")
		parentFolderIndex := slices.Index(splitString, name)

		// If parent folder found but with a different orgID, continue
		if parentFolderIndex != -1 && node.OrgId != orgID {
			foundFolderDifferentOrg = true
			continue
			// If parent folder found with the correct orgID
		} else if parentFolderIndex != -1 && node.OrgId == orgID {
			foundFolderWithOrg = true
		}

		childFolderIndex := slices.Index(splitString, node.Name)
		if parentFolderIndex < childFolderIndex && parentFolderIndex != -1 {
			childFolders = append(childFolders, node)
		}
	}

	// Check if parent folder is never found
	if !foundFolderDifferentOrg && !foundFolderWithOrg {
		err = errors.New("Folder does not exist")
		return childFolders, err
	}

	// Check if parent folder is found with different org
	if foundFolderDifferentOrg && len(childFolders) == 0 {
		err = errors.New("Folder does not exist in the specified organization")
		return childFolders, err
	}

	return childFolders, nil
}
