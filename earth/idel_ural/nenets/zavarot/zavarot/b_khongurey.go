package zavarot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洪古列伊KhongureyBarony struct {
	feud.BaseBarony
}

var BKhongurey洪古列伊 feud.Barony = &洪古列伊KhongureyBarony{}

func init() {
    f := BKhongurey洪古列伊.(*洪古列伊KhongureyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khongurey",
		TitleName: "洪古列伊",
		TitleCode: "b_khongurey",
	}
}
