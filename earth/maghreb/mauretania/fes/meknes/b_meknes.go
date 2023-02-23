package meknes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅克内斯MeknesBarony struct {
	feud.BaseBarony
}

var BMeknes梅克内斯 feud.Barony = &梅克内斯MeknesBarony{}

func init() {
	f := BMeknes梅克内斯.(*梅克内斯MeknesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "meknes",
		TitleName: "梅克内斯",
		TitleCode: "b_meknes",
	}
}
