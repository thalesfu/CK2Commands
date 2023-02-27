package monreal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比勒达BildadBarony struct {
	feud.BaseBarony
}

var BBildad比勒达 feud.Barony = &比勒达BildadBarony{}

func init() {
    f := BBildad比勒达.(*比勒达BildadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bildad",
		TitleName: "比勒达",
		TitleCode: "b_bildad",
	}
}
