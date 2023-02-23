package chur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈恩费尔德MaienfeldBarony struct {
	feud.BaseBarony
}

var BMaienfeld迈恩费尔德 feud.Barony = &迈恩费尔德MaienfeldBarony{}

func init() {
	f := BMaienfeld迈恩费尔德.(*迈恩费尔德MaienfeldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maienfeld",
		TitleName: "迈恩费尔德",
		TitleCode: "b_maienfeld",
	}
}
