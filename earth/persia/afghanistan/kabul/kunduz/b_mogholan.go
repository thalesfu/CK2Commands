package kunduz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙哥兰MogholanBarony struct {
	feud.BaseBarony
}

var BMogholan蒙哥兰 feud.Barony = &蒙哥兰MogholanBarony{}

func init() {
    f := BMogholan蒙哥兰.(*蒙哥兰MogholanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mogholan",
		TitleName: "蒙哥兰",
		TitleCode: "b_mogholan",
	}
}
