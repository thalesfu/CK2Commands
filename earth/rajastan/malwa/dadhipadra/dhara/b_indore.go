package dhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 印多尔IndoreBarony struct {
	feud.BaseBarony
}

var BIndore印多尔 feud.Barony = &印多尔IndoreBarony{}

func init() {
	f := BIndore印多尔.(*印多尔IndoreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "indore",
		TitleName: "印多尔",
		TitleCode: "b_indore",
	}
}
