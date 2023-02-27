package teluch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩勒PerreBarony struct {
	feud.BaseBarony
}

var BPerre佩勒 feud.Barony = &佩勒PerreBarony{}

func init() {
    f := BPerre佩勒.(*佩勒PerreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "perre",
		TitleName: "佩勒",
		TitleCode: "b_perre",
	}
}
