package amman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富海斯FuheisBarony struct {
	feud.BaseBarony
}

var BFuheis富海斯 feud.Barony = &富海斯FuheisBarony{}

func init() {
	f := BFuheis富海斯.(*富海斯FuheisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fuheis",
		TitleName: "富海斯",
		TitleCode: "b_fuheis",
	}
}
