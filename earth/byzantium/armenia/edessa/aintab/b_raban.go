package aintab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷班RabanBarony struct {
	feud.BaseBarony
}

var BRaban雷班 feud.Barony = &雷班RabanBarony{}

func init() {
	f := BRaban雷班.(*雷班RabanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raban",
		TitleName: "雷班",
		TitleCode: "b_raban",
	}
}
