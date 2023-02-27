package starodub

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泽列纳波利亚纳Zelena_polyanaBarony struct {
	feud.BaseBarony
}

var BZelena_polyana泽列纳波利亚纳 feud.Barony = &泽列纳波利亚纳Zelena_polyanaBarony{}

func init() {
    f := BZelena_polyana泽列纳波利亚纳.(*泽列纳波利亚纳Zelena_polyanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zelena_polyana",
		TitleName: "泽列纳波利亚纳",
		TitleCode: "b_zelena_polyana",
	}
}
