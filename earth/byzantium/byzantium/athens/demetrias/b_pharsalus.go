package demetrias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法尔萨卢斯PharsalusBarony struct {
	feud.BaseBarony
}

var BPharsalus法尔萨卢斯 feud.Barony = &法尔萨卢斯PharsalusBarony{}

func init() {
    f := BPharsalus法尔萨卢斯.(*法尔萨卢斯PharsalusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pharsalus",
		TitleName: "法尔萨卢斯",
		TitleCode: "b_pharsalus",
	}
}
