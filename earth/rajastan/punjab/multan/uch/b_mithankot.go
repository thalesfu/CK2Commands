package uch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米坦科特MithankotBarony struct {
	feud.BaseBarony
}

var BMithankot米坦科特 feud.Barony = &米坦科特MithankotBarony{}

func init() {
	f := BMithankot米坦科特.(*米坦科特MithankotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mithankot",
		TitleName: "米坦科特",
		TitleCode: "b_mithankot",
	}
}
