package gent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多尔尼克DoornikBarony struct {
	feud.BaseBarony
}

var BDoornik多尔尼克 feud.Barony = &多尔尼克DoornikBarony{}

func init() {
	f := BDoornik多尔尼克.(*多尔尼克DoornikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "doornik",
		TitleName: "多尔尼克",
		TitleCode: "b_doornik",
	}
}
