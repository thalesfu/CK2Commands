package dhamalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩尔比MorviBarony struct {
	feud.BaseBarony
}

var BMorvi摩尔比 feud.Barony = &摩尔比MorviBarony{}

func init() {
    f := BMorvi摩尔比.(*摩尔比MorviBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "morvi",
		TitleName: "摩尔比",
		TitleCode: "b_morvi",
	}
}
