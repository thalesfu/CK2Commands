package cornwall

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朗斯顿LauncestonBarony struct {
	feud.BaseBarony
}

var BLaunceston朗斯顿 feud.Barony = &朗斯顿LauncestonBarony{}

func init() {
	f := BLaunceston朗斯顿.(*朗斯顿LauncestonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "launceston",
		TitleName: "朗斯顿",
		TitleCode: "b_launceston",
	}
}
