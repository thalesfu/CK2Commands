package tui

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙布雷ShabryBarony struct {
	feud.BaseBarony
}

var BShabry沙布雷 feud.Barony = &沙布雷ShabryBarony{}

func init() {
    f := BShabry沙布雷.(*沙布雷ShabryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shabry",
		TitleName: "沙布雷",
		TitleCode: "b_shabry",
	}
}
