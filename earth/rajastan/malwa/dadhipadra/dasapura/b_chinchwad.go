package dasapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 真遮婆荼ChinchwadBarony struct {
	feud.BaseBarony
}

var BChinchwad真遮婆荼 feud.Barony = &真遮婆荼ChinchwadBarony{}

func init() {
	f := BChinchwad真遮婆荼.(*真遮婆荼ChinchwadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chinchwad",
		TitleName: "真遮婆荼",
		TitleCode: "b_chinchwad",
	}
}
