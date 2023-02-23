package kexholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶卡伯格JekaborgBarony struct {
	feud.BaseBarony
}

var BJekaborg耶卡伯格 feud.Barony = &耶卡伯格JekaborgBarony{}

func init() {
	f := BJekaborg耶卡伯格.(*耶卡伯格JekaborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jekaborg",
		TitleName: "耶卡伯格",
		TitleCode: "b_jekaborg",
	}
}
