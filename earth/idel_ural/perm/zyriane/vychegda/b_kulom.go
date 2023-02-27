package vychegda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库洛姆KulomBarony struct {
	feud.BaseBarony
}

var BKulom库洛姆 feud.Barony = &库洛姆KulomBarony{}

func init() {
    f := BKulom库洛姆.(*库洛姆KulomBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kulom",
		TitleName: "库洛姆",
		TitleCode: "b_kulom",
	}
}
