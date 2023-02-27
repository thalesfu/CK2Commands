package suvarnapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗尼多补罗VinitapuraBarony struct {
	feud.BaseBarony
}

var BVinitapura毗尼多补罗 feud.Barony = &毗尼多补罗VinitapuraBarony{}

func init() {
    f := BVinitapura毗尼多补罗.(*毗尼多补罗VinitapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vinitapura",
		TitleName: "毗尼多补罗",
		TitleCode: "b_vinitapura",
	}
}
