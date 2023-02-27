package banavasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旃陀罗古提CandraguttiBarony struct {
	feud.BaseBarony
}

var BCandragutti旃陀罗古提 feud.Barony = &旃陀罗古提CandraguttiBarony{}

func init() {
    f := BCandragutti旃陀罗古提.(*旃陀罗古提CandraguttiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "candragutti",
		TitleName: "旃陀罗古提",
		TitleCode: "b_candragutti",
	}
}
