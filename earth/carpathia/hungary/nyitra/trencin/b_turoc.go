package trencin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图罗茨TurocBarony struct {
	feud.BaseBarony
}

var BTuroc图罗茨 feud.Barony = &图罗茨TurocBarony{}

func init() {
	f := BTuroc图罗茨.(*图罗茨TurocBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turoc",
		TitleName: "图罗茨",
		TitleCode: "b_turoc",
	}
}
