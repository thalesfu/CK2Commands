package luneburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃弗恩EvernBarony struct {
	feud.BaseBarony
}

var BEvern埃弗恩 feud.Barony = &埃弗恩EvernBarony{}

func init() {
    f := BEvern埃弗恩.(*埃弗恩EvernBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "evern",
		TitleName: "埃弗恩",
		TitleCode: "b_evern",
	}
}
