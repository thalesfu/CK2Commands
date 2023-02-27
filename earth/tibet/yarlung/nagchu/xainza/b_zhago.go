package xainza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 下过ZhagoBarony struct {
	feud.BaseBarony
}

var BZhago下过 feud.Barony = &下过ZhagoBarony{}

func init() {
    f := BZhago下过.(*下过ZhagoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhago",
		TitleName: "下过",
		TitleCode: "b_zhago",
	}
}
