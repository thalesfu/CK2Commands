package pettau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赛茨SeitzBarony struct {
	feud.BaseBarony
}

var BSeitz赛茨 feud.Barony = &赛茨SeitzBarony{}

func init() {
	f := BSeitz赛茨.(*赛茨SeitzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "seitz",
		TitleName: "赛茨",
		TitleCode: "b_seitz",
	}
}
