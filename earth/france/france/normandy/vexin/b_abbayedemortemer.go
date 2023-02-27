package vexin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫特梅修道院AbbayedemortemerBarony struct {
	feud.BaseBarony
}

var BAbbayedemortemer莫特梅修道院 feud.Barony = &莫特梅修道院AbbayedemortemerBarony{}

func init() {
    f := BAbbayedemortemer莫特梅修道院.(*莫特梅修道院AbbayedemortemerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abbayedemortemer",
		TitleName: "莫特梅修道院",
		TitleCode: "b_abbayedemortemer",
	}
}
