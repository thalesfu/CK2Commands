package pisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔泰拉VolterraBarony struct {
	feud.BaseBarony
}

var BVolterra沃尔泰拉 feud.Barony = &沃尔泰拉VolterraBarony{}

func init() {
	f := BVolterra沃尔泰拉.(*沃尔泰拉VolterraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "volterra",
		TitleName: "沃尔泰拉",
		TitleCode: "b_volterra",
	}
}
