package stettin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔巴茨KolbatzBarony struct {
	feud.BaseBarony
}

var BKolbatz科尔巴茨 feud.Barony = &科尔巴茨KolbatzBarony{}

func init() {
    f := BKolbatz科尔巴茨.(*科尔巴茨KolbatzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolbatz",
		TitleName: "科尔巴茨",
		TitleCode: "b_kolbatz",
	}
}
