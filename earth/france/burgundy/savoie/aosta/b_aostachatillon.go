package aosta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙蒂永AostachatillonBarony struct {
	feud.BaseBarony
}

var BAostachatillon沙蒂永 feud.Barony = &沙蒂永AostachatillonBarony{}

func init() {
    f := BAostachatillon沙蒂永.(*沙蒂永AostachatillonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aostachatillon",
		TitleName: "沙蒂永",
		TitleCode: "b_aostachatillon",
	}
}
