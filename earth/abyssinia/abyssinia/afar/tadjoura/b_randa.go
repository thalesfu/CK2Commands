package tadjoura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰达RandaBarony struct {
	feud.BaseBarony
}

var BRanda兰达 feud.Barony = &兰达RandaBarony{}

func init() {
    f := BRanda兰达.(*兰达RandaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "randa",
		TitleName: "兰达",
		TitleCode: "b_randa",
	}
}
