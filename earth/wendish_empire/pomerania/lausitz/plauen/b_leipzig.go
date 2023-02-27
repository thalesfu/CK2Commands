package plauen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱比锡LeipzigBarony struct {
	feud.BaseBarony
}

var BLeipzig莱比锡 feud.Barony = &莱比锡LeipzigBarony{}

func init() {
    f := BLeipzig莱比锡.(*莱比锡LeipzigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leipzig",
		TitleName: "莱比锡",
		TitleCode: "b_leipzig",
	}
}
