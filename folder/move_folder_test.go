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
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.delta"},
			},
			want: []folder.Folder{
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
			want: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "golf.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "golf.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "alpha.delta"},
				{Name: "golf", OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), Paths: "golf"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			got, _ := f.MoveFolder(tt.name, tt.dst)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveFolder() = \n%v, \nwant %v", got, tt.want)
			}
		})
	}
}
