package karbala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾因塔姆尔AinaltamurBarony struct {
	feud.BaseBarony
}

var BAinaltamur艾因塔姆尔 feud.Barony = &艾因塔姆尔AinaltamurBarony{}

func init() {
	f := BAinaltamur艾因塔姆尔.(*艾因塔姆尔AinaltamurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ainaltamur",
		TitleName: "艾因塔姆尔",
		TitleCode: "b_ainaltamur",
	}
}
