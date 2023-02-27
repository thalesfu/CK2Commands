package orkhon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 呼图勒KhutulBarony struct {
	feud.BaseBarony
}

var BKhutul呼图勒 feud.Barony = &呼图勒KhutulBarony{}

func init() {
    f := BKhutul呼图勒.(*呼图勒KhutulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khutul",
		TitleName: "呼图勒",
		TitleCode: "b_khutul",
	}
}
