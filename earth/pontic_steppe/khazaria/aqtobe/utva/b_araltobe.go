package utva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉尔托别AraltobeBarony struct {
	feud.BaseBarony
}

var BAraltobe阿拉尔托别 feud.Barony = &阿拉尔托别AraltobeBarony{}

func init() {
    f := BAraltobe阿拉尔托别.(*阿拉尔托别AraltobeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "araltobe",
		TitleName: "阿拉尔托别",
		TitleCode: "b_araltobe",
	}
}
