package tharasset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖勒敏GuelmimBarony struct {
	feud.BaseBarony
}

var BGuelmim盖勒敏 feud.Barony = &盖勒敏GuelmimBarony{}

func init() {
	f := BGuelmim盖勒敏.(*盖勒敏GuelmimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guelmim",
		TitleName: "盖勒敏",
		TitleCode: "b_guelmim",
	}
}
