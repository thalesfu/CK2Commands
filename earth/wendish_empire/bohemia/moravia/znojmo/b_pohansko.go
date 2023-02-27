package znojmo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波汉斯科PohanskoBarony struct {
	feud.BaseBarony
}

var BPohansko波汉斯科 feud.Barony = &波汉斯科PohanskoBarony{}

func init() {
    f := BPohansko波汉斯科.(*波汉斯科PohanskoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pohansko",
		TitleName: "波汉斯科",
		TitleCode: "b_pohansko",
	}
}
