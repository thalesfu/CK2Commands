package olvia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯塔夫基StavkyBarony struct {
	feud.BaseBarony
}

var BStavky斯塔夫基 feud.Barony = &斯塔夫基StavkyBarony{}

func init() {
    f := BStavky斯塔夫基.(*斯塔夫基StavkyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stavky",
		TitleName: "斯塔夫基",
		TitleCode: "b_stavky",
	}
}
