package mecca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图赖拜TurubahBarony struct {
	feud.BaseBarony
}

var BTurubah图赖拜 feud.Barony = &图赖拜TurubahBarony{}

func init() {
    f := BTurubah图赖拜.(*图赖拜TurubahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turubah",
		TitleName: "图赖拜",
		TitleCode: "b_turubah",
	}
}
