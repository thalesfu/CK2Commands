package ryazan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图尔拉托沃TurlatovoBarony struct {
	feud.BaseBarony
}

var BTurlatovo图尔拉托沃 feud.Barony = &图尔拉托沃TurlatovoBarony{}

func init() {
	f := BTurlatovo图尔拉托沃.(*图尔拉托沃TurlatovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turlatovo",
		TitleName: "图尔拉托沃",
		TitleCode: "b_turlatovo",
	}
}
