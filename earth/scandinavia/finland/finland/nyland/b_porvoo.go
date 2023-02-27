package nyland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波尔沃PorvooBarony struct {
	feud.BaseBarony
}

var BPorvoo波尔沃 feud.Barony = &波尔沃PorvooBarony{}

func init() {
    f := BPorvoo波尔沃.(*波尔沃PorvooBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "porvoo",
		TitleName: "波尔沃",
		TitleCode: "b_porvoo",
	}
}
