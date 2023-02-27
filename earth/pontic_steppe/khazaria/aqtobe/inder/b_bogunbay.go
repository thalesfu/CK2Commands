package inder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博贡拜BogunbayBarony struct {
	feud.BaseBarony
}

var BBogunbay博贡拜 feud.Barony = &博贡拜BogunbayBarony{}

func init() {
    f := BBogunbay博贡拜.(*博贡拜BogunbayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bogunbay",
		TitleName: "博贡拜",
		TitleCode: "b_bogunbay",
	}
}
