package chikoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赤苦ChikoiBarony struct {
	feud.BaseBarony
}

var BChikoi赤苦 feud.Barony = &赤苦ChikoiBarony{}

func init() {
	f := BChikoi赤苦.(*赤苦ChikoiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chikoi",
		TitleName: "赤苦",
		TitleCode: "b_chikoi",
	}
}
