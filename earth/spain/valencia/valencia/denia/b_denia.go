package denia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德尼亚DeniaBarony struct {
	feud.BaseBarony
}

var BDenia德尼亚 feud.Barony = &德尼亚DeniaBarony{}

func init() {
	f := BDenia德尼亚.(*德尼亚DeniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "denia",
		TitleName: "德尼亚",
		TitleCode: "b_denia",
	}
}
