package ulm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔巴赫ErbachBarony struct {
	feud.BaseBarony
}

var BErbach埃尔巴赫 feud.Barony = &埃尔巴赫ErbachBarony{}

func init() {
	f := BErbach埃尔巴赫.(*埃尔巴赫ErbachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erbach",
		TitleName: "埃尔巴赫",
		TitleCode: "b_erbach",
	}
}
