package genoa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GenoaKingdom interface {
    feud.Kingdom
}

type 热那亚GenoaKingdom struct {
	feud.BaseKingdom
}

var KGenoa热那亚 GenoaKingdom = &热那亚GenoaKingdom{}

func init() {
	f := KGenoa热那亚.(*热那亚GenoaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "genoa",
		TitleName: "热那亚",
		TitleCode: "k_genoa",
		Dukes:  map[string]feud.Duke{},
	}

}
