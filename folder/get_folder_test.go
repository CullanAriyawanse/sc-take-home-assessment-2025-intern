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
		{
			name:  "Multiple folders",
			orgID: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"),
			folders: []folder.Folder{
				{Name: "test-folder", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"), Paths: "test-folder"},
				{Name: "test-folder1", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a2"), Paths: "test-folder1"},
			},
			want: []folder.Folder{
				{Name: "test-folder", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"), Paths: "test-folder"},
			},
		},
		{
			name:  "Multiple folders same name",
			orgID: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"),
			folders: []folder.Folder{
				{Name: "test-folder", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"), Paths: "test-folder"},
				{Name: "test-folder", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a2"), Paths: "test-folder"},
			},
			want: []folder.Folder{
				{Name: "test-folder", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"), Paths: "test-folder"},
			},
		},
		{
			name:  "Invalid OrgId",
			orgID: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a2"),
			folders: []folder.Folder{
				{Name: "test-folder", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"), Paths: "test-folder"},
			},
			want: []folder.Folder{},
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

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name             string
		orgID            uuid.UUID
		parentFolderName string
		folders          []folder.Folder
		expected         []folder.Folder
		expectedError    string
	}{
		{
			name:             "Base Case",
			orgID:            uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			parentFolderName: "creative-scalphunter",
			folders: []folder.Folder{
				{Name: "test-folder", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "test-folder"},
				{Name: "creative-scalphunter", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "creative-scalphunter"},
				{Name: "clear-arclight", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "creative-scalphunter.clear-arclight"},
			},
			expected: []folder.Folder{
				{Name: "clear-arclight", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "creative-scalphunter.clear-arclight"},
			},
		},
		{
			name:             "No child folders",
			orgID:            uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"),
			parentFolderName: "test-folder",
			folders: []folder.Folder{
				{Name: "test-folder", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"), Paths: "test-folder"},
			},
			expected: []folder.Folder{},
		},
		{
			name:             "Same folder name different orgId",
			orgID:            uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"),
			parentFolderName: "test-folder",
			folders: []folder.Folder{
				{Name: "test-folder", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"), Paths: "test-folder"},
				{Name: "folder1", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"), Paths: "test-folder.folder1"},
				{Name: "folder1", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-ffffffffffff"), Paths: "test-folder.folder1"}, // Different org
			},
			expected: []folder.Folder{
				{Name: "folder1", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"), Paths: "test-folder.folder1"},
			},
		},
		{
			name:             "Invalid OrgID",
			orgID:            uuid.FromStringOrNil(""),
			parentFolderName: "test-folder",
			folders: []folder.Folder{
				{Name: "test-folder", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"), Paths: "test-folder"},
			},
			expected: []folder.Folder{},
		},
		{
			name:             "Invalid folder name",
			orgID:            uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"),
			parentFolderName: "",
			folders: []folder.Folder{
				{Name: "test-folder", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"), Paths: "test-folder"},
			},
			expectedError: "Folder does not exist",
		},
		{
			name:             "Invalid folder name for given orgID",
			orgID:            uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-fffffffffff"),
			parentFolderName: "test-folder",
			folders: []folder.Folder{
				{Name: "test-folder", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a1"), Paths: "test-folder"},
			},
			expectedError: "Folder does not exist in the specified organization",
		},
		{
			name:             "Multiple subfolders",
			orgID:            uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			parentFolderName: "creative-scalphunter",
			folders: []folder.Folder{
				{Name: "creative-scalphunter", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "creative-scalphunter"},
				{Name: "clear-arclight", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "creative-scalphunter.clear-arclight"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "creative-scalphunter.clear-arclight.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "creative-scalphunter.clear-arclight.charlie"},
			},
			expected: []folder.Folder{
				{Name: "clear-arclight", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "creative-scalphunter.clear-arclight"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "creative-scalphunter.clear-arclight.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "creative-scalphunter.clear-arclight.charlie"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			got, err := f.GetAllChildFolders(tt.orgID, tt.parentFolderName)
			if tt.expectedError != "" {
				if err == nil || err.Error() != tt.expectedError {
					t.Errorf("GetAllChildFolders() error = %v, expected error %v", err, tt.expectedError)
				}
			} else {
				if !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("GetAllChildFolders() = %v, expected %v", got, tt.expected)
				}
			}
		})
	}
}
