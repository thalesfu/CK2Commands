package tabaristan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲鲁兹库赫FiruzkuhBarony struct {
	feud.BaseBarony
}

var BFiruzkuh菲鲁兹库赫 feud.Barony = &菲鲁兹库赫FiruzkuhBarony{}

func init() {
	f := BFiruzkuh菲鲁兹库赫.(*菲鲁兹库赫FiruzkuhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "firuzkuh",
		TitleName: "菲鲁兹库赫",
		TitleCode: "b_firuzkuh",
	}
}
