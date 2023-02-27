package shish

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希什ShishBarony struct {
	feud.BaseBarony
}

var BShish希什 feud.Barony = &希什ShishBarony{}

func init() {
    f := BShish希什.(*希什ShishBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shish",
		TitleName: "希什",
		TitleCode: "b_shish",
	}
}
