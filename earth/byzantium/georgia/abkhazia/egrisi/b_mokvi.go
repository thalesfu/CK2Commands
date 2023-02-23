package egrisi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫克维MokviBarony struct {
	feud.BaseBarony
}

var BMokvi莫克维 feud.Barony = &莫克维MokviBarony{}

func init() {
	f := BMokvi莫克维.(*莫克维MokviBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mokvi",
		TitleName: "莫克维",
		TitleCode: "b_mokvi",
	}
}
