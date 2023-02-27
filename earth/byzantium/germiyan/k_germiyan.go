package germiyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GermiyanKingdom interface {
    feud.Kingdom
}

type 格尔米扬GermiyanKingdom struct {
	feud.BaseKingdom
}

var KGermiyan格尔米扬 GermiyanKingdom = &格尔米扬GermiyanKingdom{}

func init() {
	f := KGermiyan格尔米扬.(*格尔米扬GermiyanKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "germiyan",
		TitleName: "格尔米扬",
		TitleCode: "k_germiyan",
		Dukes:  map[string]feud.Duke{},
	}

}
