package garhwal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提伐迦DevalgharBarony struct {
	feud.BaseBarony
}

var BDevalghar提伐迦 feud.Barony = &提伐迦DevalgharBarony{}

func init() {
    f := BDevalghar提伐迦.(*提伐迦DevalgharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "devalghar",
		TitleName: "提伐迦",
		TitleCode: "b_devalghar",
	}
}
