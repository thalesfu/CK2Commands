package zvenyhorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗戈任RogozhynBarony struct {
	feud.BaseBarony
}

var BRogozhyn罗戈任 feud.Barony = &罗戈任RogozhynBarony{}

func init() {
    f := BRogozhyn罗戈任.(*罗戈任RogozhynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rogozhyn",
		TitleName: "罗戈任",
		TitleCode: "b_rogozhyn",
	}
}
