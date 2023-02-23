package albret

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗克福尔RoquefortBarony struct {
	feud.BaseBarony
}

var BRoquefort罗克福尔 feud.Barony = &罗克福尔RoquefortBarony{}

func init() {
	f := BRoquefort罗克福尔.(*罗克福尔RoquefortBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roquefort",
		TitleName: "罗克福尔",
		TitleCode: "b_roquefort",
	}
}
