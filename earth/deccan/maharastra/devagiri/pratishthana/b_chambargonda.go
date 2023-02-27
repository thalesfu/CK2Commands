package pratishthana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 占跋罗郡荼ChambargondaBarony struct {
	feud.BaseBarony
}

var BChambargonda占跋罗郡荼 feud.Barony = &占跋罗郡荼ChambargondaBarony{}

func init() {
    f := BChambargonda占跋罗郡荼.(*占跋罗郡荼ChambargondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chambargonda",
		TitleName: "占跋罗郡荼",
		TitleCode: "b_chambargonda",
	}
}
