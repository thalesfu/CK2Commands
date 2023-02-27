package chunar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗荼提BidadiBarony struct {
	feud.BaseBarony
}

var BBidadi毗荼提 feud.Barony = &毗荼提BidadiBarony{}

func init() {
    f := BBidadi毗荼提.(*毗荼提BidadiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bidadi",
		TitleName: "毗荼提",
		TitleCode: "b_bidadi",
	}
}
