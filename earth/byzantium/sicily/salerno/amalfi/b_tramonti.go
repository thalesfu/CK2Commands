package amalfi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特拉蒙蒂TramontiBarony struct {
	feud.BaseBarony
}

var BTramonti特拉蒙蒂 feud.Barony = &特拉蒙蒂TramontiBarony{}

func init() {
	f := BTramonti特拉蒙蒂.(*特拉蒙蒂TramontiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tramonti",
		TitleName: "特拉蒙蒂",
		TitleCode: "b_tramonti",
	}
}
