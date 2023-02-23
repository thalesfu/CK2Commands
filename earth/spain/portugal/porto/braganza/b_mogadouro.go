package braganza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫加多鲁MogadouroBarony struct {
	feud.BaseBarony
}

var BMogadouro莫加多鲁 feud.Barony = &莫加多鲁MogadouroBarony{}

func init() {
	f := BMogadouro莫加多鲁.(*莫加多鲁MogadouroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mogadouro",
		TitleName: "莫加多鲁",
		TitleCode: "b_mogadouro",
	}
}
