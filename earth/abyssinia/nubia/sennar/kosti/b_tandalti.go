package kosti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坦达勒提TandaltiBarony struct {
	feud.BaseBarony
}

var BTandalti坦达勒提 feud.Barony = &坦达勒提TandaltiBarony{}

func init() {
    f := BTandalti坦达勒提.(*坦达勒提TandaltiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tandalti",
		TitleName: "坦达勒提",
		TitleCode: "b_tandalti",
	}
}
