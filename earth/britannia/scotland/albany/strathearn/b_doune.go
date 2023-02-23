package strathearn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜恩DouneBarony struct {
	feud.BaseBarony
}

var BDoune杜恩 feud.Barony = &杜恩DouneBarony{}

func init() {
	f := BDoune杜恩.(*杜恩DouneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "doune",
		TitleName: "杜恩",
		TitleCode: "b_doune",
	}
}
