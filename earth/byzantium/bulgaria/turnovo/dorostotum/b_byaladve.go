package dorostotum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比亚拉德韦ByaladveBarony struct {
	feud.BaseBarony
}

var BByaladve比亚拉德韦 feud.Barony = &比亚拉德韦ByaladveBarony{}

func init() {
	f := BByaladve比亚拉德韦.(*比亚拉德韦ByaladveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "byaladve",
		TitleName: "比亚拉德韦",
		TitleCode: "b_byaladve",
	}
}
