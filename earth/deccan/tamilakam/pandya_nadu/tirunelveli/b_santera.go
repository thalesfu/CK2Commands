package tirunelveli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑谛罗SanteraBarony struct {
	feud.BaseBarony
}

var BSantera桑谛罗 feud.Barony = &桑谛罗SanteraBarony{}

func init() {
    f := BSantera桑谛罗.(*桑谛罗SanteraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "santera",
		TitleName: "桑谛罗",
		TitleCode: "b_santera",
	}
}
