package brugge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯勒伊斯SluysBarony struct {
	feud.BaseBarony
}

var BSluys斯勒伊斯 feud.Barony = &斯勒伊斯SluysBarony{}

func init() {
    f := BSluys斯勒伊斯.(*斯勒伊斯SluysBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sluys",
		TitleName: "斯勒伊斯",
		TitleCode: "b_sluys",
	}
}
