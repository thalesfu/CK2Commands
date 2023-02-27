package seleukeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达利桑佐斯DalisandusBarony struct {
	feud.BaseBarony
}

var BDalisandus达利桑佐斯 feud.Barony = &达利桑佐斯DalisandusBarony{}

func init() {
    f := BDalisandus达利桑佐斯.(*达利桑佐斯DalisandusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dalisandus",
		TitleName: "达利桑佐斯",
		TitleCode: "b_dalisandus",
	}
}
