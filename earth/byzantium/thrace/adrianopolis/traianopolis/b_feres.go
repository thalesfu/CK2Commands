package traianopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费雷斯FeresBarony struct {
	feud.BaseBarony
}

var BFeres费雷斯 feud.Barony = &费雷斯FeresBarony{}

func init() {
	f := BFeres费雷斯.(*费雷斯FeresBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "feres",
		TitleName: "费雷斯",
		TitleCode: "b_feres",
	}
}
