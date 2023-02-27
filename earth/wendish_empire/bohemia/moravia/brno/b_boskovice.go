package brno

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博斯科维采BoskoviceBarony struct {
	feud.BaseBarony
}

var BBoskovice博斯科维采 feud.Barony = &博斯科维采BoskoviceBarony{}

func init() {
    f := BBoskovice博斯科维采.(*博斯科维采BoskoviceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boskovice",
		TitleName: "博斯科维采",
		TitleCode: "b_boskovice",
	}
}
