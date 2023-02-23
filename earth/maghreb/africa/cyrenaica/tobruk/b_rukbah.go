package tobruk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁克班RukbahBarony struct {
	feud.BaseBarony
}

var BRukbah鲁克班 feud.Barony = &鲁克班RukbahBarony{}

func init() {
	f := BRukbah鲁克班.(*鲁克班RukbahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rukbah",
		TitleName: "鲁克班",
		TitleCode: "b_rukbah",
	}
}
