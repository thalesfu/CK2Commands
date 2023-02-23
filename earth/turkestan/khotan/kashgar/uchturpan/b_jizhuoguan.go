package uchturpan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 济浊馆JizhuoguanBarony struct {
	feud.BaseBarony
}

var BJizhuoguan济浊馆 feud.Barony = &济浊馆JizhuoguanBarony{}

func init() {
	f := BJizhuoguan济浊馆.(*济浊馆JizhuoguanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jizhuoguan",
		TitleName: "济浊馆",
		TitleCode: "b_jizhuoguan",
	}
}
