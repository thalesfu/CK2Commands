package hradyzk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波尔塔瓦PoltavaBarony struct {
	feud.BaseBarony
}

var BPoltava波尔塔瓦 feud.Barony = &波尔塔瓦PoltavaBarony{}

func init() {
	f := BPoltava波尔塔瓦.(*波尔塔瓦PoltavaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "poltava",
		TitleName: "波尔塔瓦",
		TitleCode: "b_poltava",
	}
}
