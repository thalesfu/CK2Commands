package skane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫尔辛堡HelsingborgBarony struct {
	feud.BaseBarony
}

var BHelsingborg赫尔辛堡 feud.Barony = &赫尔辛堡HelsingborgBarony{}

func init() {
    f := BHelsingborg赫尔辛堡.(*赫尔辛堡HelsingborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "helsingborg",
		TitleName: "赫尔辛堡",
		TitleCode: "b_helsingborg",
	}
}
