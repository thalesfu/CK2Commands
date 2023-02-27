package savoie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托农ThononBarony struct {
	feud.BaseBarony
}

var BThonon托农 feud.Barony = &托农ThononBarony{}

func init() {
    f := BThonon托农.(*托农ThononBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thonon",
		TitleName: "托农",
		TitleCode: "b_thonon",
	}
}
