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
		orgID         uuid.UUID
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
			testName: "Base Case 2",
			name:     "bravo",
			dst:      "golf",
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.delta"},
				{Name: "golf", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "golf"},
			},
			expected: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "golf.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "golf.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.delta"},
				{Name: "golf", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "golf"},
			},
		},
		{
			testName: "Error moving folder to child of itself",
			name:     "bravo",
			dst:      "charlie",
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.bravo.charlie"},
			},
			expectedError: "Error: Cannot move a folder to a child of itself",
		},
		{
			testName: "Error moving folder to itself",
			name:     "bravo",
			dst:      "bravo",
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.bravo.charlie"},
			},
			expectedError: "Error: Cannot move a folder to itself",
		},
		{
			testName: "Error moving folder to a different organisation",
			name:     "bravo",
			dst:      "bravo",
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.bravo.charlie"},
			},
			expectedError: "Error: Cannot move a folder to a different organization",
		},
		{
			testName: "Error moving folder that does not exist",
			name:     "bravo",
			dst:      "bravo",
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.bravo.charlie"},
			},
			expectedError: "Error: Source folder does not exist",
		},
		{
			testName: "Error moving folder to destination folder that does not exist",
			name:     "bravo",
			dst:      "bravo",
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.bravo.charlie"},
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
