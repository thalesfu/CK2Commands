package monyul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达旺TawangBarony struct {
	feud.BaseBarony
}

var BTawang达旺 feud.Barony = &达旺TawangBarony{}

func init() {
	f := BTawang达旺.(*达旺TawangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tawang",
		TitleName: "达旺",
		TitleCode: "b_tawang",
	}
}
