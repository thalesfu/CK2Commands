package dwarasamudra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏萨乌尔SosavurBarony struct {
	feud.BaseBarony
}

var BSosavur苏萨乌尔 feud.Barony = &苏萨乌尔SosavurBarony{}

func init() {
    f := BSosavur苏萨乌尔.(*苏萨乌尔SosavurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sosavur",
		TitleName: "苏萨乌尔",
		TitleCode: "b_sosavur",
	}
}
