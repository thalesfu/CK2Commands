package tenkasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丁迪古尔DindigulBarony struct {
	feud.BaseBarony
}

var BDindigul丁迪古尔 feud.Barony = &丁迪古尔DindigulBarony{}

func init() {
    f := BDindigul丁迪古尔.(*丁迪古尔DindigulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dindigul",
		TitleName: "丁迪古尔",
		TitleCode: "b_dindigul",
	}
}
