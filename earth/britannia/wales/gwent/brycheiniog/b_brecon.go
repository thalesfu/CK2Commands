package brycheiniog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布雷肯BreconBarony struct {
	feud.BaseBarony
}

var BBrecon布雷肯 feud.Barony = &布雷肯BreconBarony{}

func init() {
    f := BBrecon布雷肯.(*布雷肯BreconBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brecon",
		TitleName: "布雷肯",
		TitleCode: "b_brecon",
	}
}
