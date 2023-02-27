package blekinge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 龙讷比RonnebyBarony struct {
	feud.BaseBarony
}

var BRonneby龙讷比 feud.Barony = &龙讷比RonnebyBarony{}

func init() {
    f := BRonneby龙讷比.(*龙讷比RonnebyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ronneby",
		TitleName: "龙讷比",
		TitleCode: "b_ronneby",
	}
}
