package vairagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艺卢KelodBarony struct {
	feud.BaseBarony
}

var BKelod艺卢 feud.Barony = &艺卢KelodBarony{}

func init() {
	f := BKelod艺卢.(*艺卢KelodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kelod",
		TitleName: "艺卢",
		TitleCode: "b_kelod",
	}
}
