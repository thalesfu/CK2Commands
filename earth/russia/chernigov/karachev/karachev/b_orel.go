package karachev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥廖尔OrelBarony struct {
	feud.BaseBarony
}

var BOrel奥廖尔 feud.Barony = &奥廖尔OrelBarony{}

func init() {
	f := BOrel奥廖尔.(*奥廖尔OrelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orel",
		TitleName: "奥廖尔",
		TitleCode: "b_orel",
	}
}
