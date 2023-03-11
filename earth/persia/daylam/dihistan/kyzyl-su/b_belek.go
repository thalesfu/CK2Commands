package kyzyl-su

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别列克BelekBarony struct {
	feud.BaseBarony
}

var BBelek别列克 feud.Barony = &别列克BelekBarony{}

func init() {
    f := BBelek别列克.(*别列克BelekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belek",
		TitleName: "别列克",
		TitleCode: "b_belek",
	}
}
