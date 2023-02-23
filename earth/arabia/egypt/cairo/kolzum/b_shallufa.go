package kolzum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莎鲁法ShallufaBarony struct {
	feud.BaseBarony
}

var BShallufa莎鲁法 feud.Barony = &莎鲁法ShallufaBarony{}

func init() {
	f := BShallufa莎鲁法.(*莎鲁法ShallufaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shallufa",
		TitleName: "莎鲁法",
		TitleCode: "b_shallufa",
	}
}
