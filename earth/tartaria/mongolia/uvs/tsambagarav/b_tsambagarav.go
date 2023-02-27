package tsambagarav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 查木巴嘎拉布TsambagaravBarony struct {
	feud.BaseBarony
}

var BTsambagarav查木巴嘎拉布 feud.Barony = &查木巴嘎拉布TsambagaravBarony{}

func init() {
    f := BTsambagarav查木巴嘎拉布.(*查木巴嘎拉布TsambagaravBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsambagarav",
		TitleName: "查木巴嘎拉布",
		TitleCode: "b_tsambagarav",
	}
}
