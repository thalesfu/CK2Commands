package lhatse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 江蚌摩齐GyangbumocheBarony struct {
	feud.BaseBarony
}

var BGyangbumoche江蚌摩齐 feud.Barony = &江蚌摩齐GyangbumocheBarony{}

func init() {
	f := BGyangbumoche江蚌摩齐.(*江蚌摩齐GyangbumocheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gyangbumoche",
		TitleName: "江蚌摩齐",
		TitleCode: "b_gyangbumoche",
	}
}
