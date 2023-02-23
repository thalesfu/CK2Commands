package cairo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫勒万HelwanBarony struct {
	feud.BaseBarony
}

var BHelwan赫勒万 feud.Barony = &赫勒万HelwanBarony{}

func init() {
	f := BHelwan赫勒万.(*赫勒万HelwanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "helwan",
		TitleName: "赫勒万",
		TitleCode: "b_helwan",
	}
}
