package tlemcen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆阿斯凯尔MuaskarBarony struct {
	feud.BaseBarony
}

var BMuaskar穆阿斯凯尔 feud.Barony = &穆阿斯凯尔MuaskarBarony{}

func init() {
    f := BMuaskar穆阿斯凯尔.(*穆阿斯凯尔MuaskarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muaskar",
		TitleName: "穆阿斯凯尔",
		TitleCode: "b_muaskar",
	}
}
