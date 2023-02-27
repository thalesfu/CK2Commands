package taghaza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅达塔MetartaBarony struct {
	feud.BaseBarony
}

var BMetarta梅达塔 feud.Barony = &梅达塔MetartaBarony{}

func init() {
    f := BMetarta梅达塔.(*梅达塔MetartaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "metarta",
		TitleName: "梅达塔",
		TitleCode: "b_metarta",
	}
}
