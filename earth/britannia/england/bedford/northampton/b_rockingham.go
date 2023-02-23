package northampton

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗金厄姆RockinghamBarony struct {
	feud.BaseBarony
}

var BRockingham罗金厄姆 feud.Barony = &罗金厄姆RockinghamBarony{}

func init() {
	f := BRockingham罗金厄姆.(*罗金厄姆RockinghamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rockingham",
		TitleName: "罗金厄姆",
		TitleCode: "b_rockingham",
	}
}
