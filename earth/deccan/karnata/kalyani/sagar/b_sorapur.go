package sagar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏拉补罗SorapurBarony struct {
	feud.BaseBarony
}

var BSorapur苏拉补罗 feud.Barony = &苏拉补罗SorapurBarony{}

func init() {
	f := BSorapur苏拉补罗.(*苏拉补罗SorapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sorapur",
		TitleName: "苏拉补罗",
		TitleCode: "b_sorapur",
	}
}
