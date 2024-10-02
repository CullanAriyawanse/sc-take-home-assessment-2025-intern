package folder

import (
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

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	folderData := GetSampleData()

	childFolders := []Folder{}

	// Split path into slice, if index of parent folder is less than child folder, add child folder to slice
	for _, node := range folderData {
		splitString := strings.Split(node.Paths, ".")
		parentFolderIndex := slices.Index(splitString, name)
		childFolderIndex := slices.Index(splitString, node.Name)
		if parentFolderIndex < childFolderIndex && parentFolderIndex != -1 {
			childFolders = append(childFolders, node)
		}
	}

	return childFolders
}
