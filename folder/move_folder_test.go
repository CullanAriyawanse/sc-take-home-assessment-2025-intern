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
		testName string
		name     string
		dst      string
		orgID    uuid.UUID
		folders  []folder.Folder
		want     []folder.Folder
	}{
		{
			testName: "Base Case",
			name:     "bravo",
			dst:      "delta",
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("org1"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("org1"), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("org1"), Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil("org1"), Paths: "alpha.delta"},
			},
			want: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("org1"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("org1"), Paths: "alpha.delta.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("org1"), Paths: "alpha.delta.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil("org1"), Paths: "alpha.delta"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			got := f.GetFoldersByOrgID(tt.orgID)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}
