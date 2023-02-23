package orava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利普托夫斯基米库拉什LiptovskymikulasBarony struct {
	feud.BaseBarony
}

var BLiptovskymikulas利普托夫斯基米库拉什 feud.Barony = &利普托夫斯基米库拉什LiptovskymikulasBarony{}

func init() {
	f := BLiptovskymikulas利普托夫斯基米库拉什.(*利普托夫斯基米库拉什LiptovskymikulasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "liptovskymikulas",
		TitleName: "利普托夫斯基米库拉什",
		TitleCode: "b_liptovskymikulas",
	}
}
