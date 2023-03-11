package cherven

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔文ChervenBarony struct {
	feud.BaseBarony
}

var BCherven切尔文 feud.Barony = &切尔文ChervenBarony{}

func init() {
    f := BCherven切尔文.(*切尔文ChervenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cherven",
		TitleName: "切尔文",
		TitleCode: "b_cherven",
	}
}
