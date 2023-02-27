package lyncestis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫拉克利亚Heraclea_lyncestisBarony struct {
	feud.BaseBarony
}

var BHeraclea_lyncestis赫拉克利亚 feud.Barony = &赫拉克利亚Heraclea_lyncestisBarony{}

func init() {
    f := BHeraclea_lyncestis赫拉克利亚.(*赫拉克利亚Heraclea_lyncestisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "heraclea_lyncestis",
		TitleName: "赫拉克利亚",
		TitleCode: "b_heraclea_lyncestis",
	}
}
