package folder_test

import (
	"reflect"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		testName      string
		name          string
		dst           string
		folders       []folder.Folder
		expected      []folder.Folder
		expectedError string
	}{
		{
			testName: "Base Case",
			name:     "bravo",
			dst:      "delta",
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.delta"},
			},
			expected: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.delta.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.delta.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.delta"},
			},
		},
		{
			testName: "Move Folder with Subfolders",
			name:     "alpha",
			dst:      "delta",
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "delta"},
			},
			expected: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "delta.alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "delta.alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "delta.alpha.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "delta"},
			},
		},
		{
			testName: "Moving root folder",
			name:     "alpha",
			dst:      "bravo",
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.delta"},
			},
			expected: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "bravo.alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "bravo.alpha.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "bravo.alpha.delta"},
			},
		},
		{
			testName: "Error moving folder to child of itself",
			name:     "alpha",
			dst:      "bravo",
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.bravo"},
			},
			expectedError: "Error: Cannot move a folder to a child of itself",
		},
		{
			testName: "Error moving folder to itself",
			name:     "alpha",
			dst:      "alpha",
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha"},
			},
			expectedError: "Error: Cannot move a folder to itself",
		},
		{
			testName: "Error moving folder to a different organisation",
			name:     "alpha",
			dst:      "bravo",
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-ffffffffffff"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "bravo"},
			},
			expectedError: "Error: Cannot move a folder to a different organization",
		},
		{
			testName: "Error moving folder that does not exist",
			name:     "invalid_folder",
			dst:      "alpha",
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha"},
			},
			expectedError: "Error: Source folder does not exist",
		},
		{
			testName: "Error moving folder to destination folder that does not exist",
			name:     "alpha",
			dst:      "invalid_folder",
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha"},
			},
			expectedError: "Error: Destination folder does not exist",
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			got, err := f.MoveFolder(tt.name, tt.dst)
			if tt.expectedError != "" {
				if err == nil || err.Error() != tt.expectedError {
					t.Errorf("moveFolder() error = %v, expected error %v", err, tt.expectedError)
				}
			} else {
				if !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("moveFolder() = \n%v, \nexpected %v", got, tt.expected)
				}
			}
		})
	}
}
