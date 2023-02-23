package manyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基科利KikkeriBarony struct {
	feud.BaseBarony
}

var BKikkeri基科利 feud.Barony = &基科利KikkeriBarony{}

func init() {
	f := BKikkeri基科利.(*基科利KikkeriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kikkeri",
		TitleName: "基科利",
		TitleCode: "b_kikkeri",
	}
}
