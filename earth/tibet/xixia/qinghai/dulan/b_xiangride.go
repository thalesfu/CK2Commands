package dulan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 香日德XiangrideBarony struct {
	feud.BaseBarony
}

var BXiangride香日德 feud.Barony = &香日德XiangrideBarony{}

func init() {
    f := BXiangride香日德.(*香日德XiangrideBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xiangride",
		TitleName: "香日德",
		TitleCode: "b_xiangride",
	}
}
