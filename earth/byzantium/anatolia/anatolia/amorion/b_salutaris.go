package amorion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨卢塔里斯SalutarisBarony struct {
	feud.BaseBarony
}

var BSalutaris萨卢塔里斯 feud.Barony = &萨卢塔里斯SalutarisBarony{}

func init() {
	f := BSalutaris萨卢塔里斯.(*萨卢塔里斯SalutarisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salutaris",
		TitleName: "萨卢塔里斯",
		TitleCode: "b_salutaris",
	}
}
