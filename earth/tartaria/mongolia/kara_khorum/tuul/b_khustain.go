package tuul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 呼斯坦KhustainBarony struct {
	feud.BaseBarony
}

var BKhustain呼斯坦 feud.Barony = &呼斯坦KhustainBarony{}

func init() {
    f := BKhustain呼斯坦.(*呼斯坦KhustainBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khustain",
		TitleName: "呼斯坦",
		TitleCode: "b_khustain",
	}
}
