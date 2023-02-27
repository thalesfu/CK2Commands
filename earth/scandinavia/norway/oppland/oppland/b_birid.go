package oppland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比里德BiridBarony struct {
	feud.BaseBarony
}

var BBirid比里德 feud.Barony = &比里德BiridBarony{}

func init() {
    f := BBirid比里德.(*比里德BiridBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "birid",
		TitleName: "比里德",
		TitleCode: "b_birid",
	}
}
