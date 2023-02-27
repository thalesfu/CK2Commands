package vaspurakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凡VanBarony struct {
	feud.BaseBarony
}

var BVan凡 feud.Barony = &凡VanBarony{}

func init() {
    f := BVan凡.(*凡VanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "van",
		TitleName: "凡",
		TitleCode: "b_van",
	}
}
