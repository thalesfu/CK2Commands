package ikonion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加斯帕德里GaspadaleBarony struct {
	feud.BaseBarony
}

var BGaspadale加斯帕德里 feud.Barony = &加斯帕德里GaspadaleBarony{}

func init() {
    f := BGaspadale加斯帕德里.(*加斯帕德里GaspadaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gaspadale",
		TitleName: "加斯帕德里",
		TitleCode: "b_gaspadale",
	}
}
