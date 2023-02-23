package bornholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伦讷RonneBarony struct {
	feud.BaseBarony
}

var BRonne伦讷 feud.Barony = &伦讷RonneBarony{}

func init() {
	f := BRonne伦讷.(*伦讷RonneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ronne",
		TitleName: "伦讷",
		TitleCode: "b_ronne",
	}
}
