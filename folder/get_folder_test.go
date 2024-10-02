package folder_test

import (
	"reflect"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

func Test_folder_GetAllFolders(t *testing.T) {
	allFolders := folder.GetAllFolders()
	expectedFirstIndex := folder.Folder{
		Name:  "test-folder",
		OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"),
		Paths: "test-folder",
	}

	gotFirstIndex := allFolders[0]

	if expectedFirstIndex != gotFirstIndex {
		t.Errorf("GetAllFolders() = %v, want %v", gotFirstIndex, expectedFirstIndex)
	}

	expectedLastIndex := folder.Folder{
		Name:  "alive-tsunami",
		OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
		Paths: "steady-insect.national-screwball.sacred-lady-shiva.quick-cyber.alive-tsunami",
	}

	gotLastIndex := allFolders[len(allFolders)-1]
	if expectedLastIndex != gotLastIndex {
		t.Errorf("GetAllFolders() = %v, want %v", gotLastIndex, expectedLastIndex)
	}
}

func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name:  "Base Case",
			orgID: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"),
			folders: []folder.Folder{
				{Name: "test-folder", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"), Paths: "test-folder"},
			},
			want: []folder.Folder{
				{Name: "test-folder", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"), Paths: "test-folder"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			got := f.GetFoldersByOrgID(tt.orgID)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFoldersByOrgID() = %v, want %v", got, tt.want)
			}
		})
	}
}
