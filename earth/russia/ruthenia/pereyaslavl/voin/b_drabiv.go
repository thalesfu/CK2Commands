package voin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德拉比夫DrabivBarony struct {
	feud.BaseBarony
}

var BDrabiv德拉比夫 feud.Barony = &德拉比夫DrabivBarony{}

func init() {
    f := BDrabiv德拉比夫.(*德拉比夫DrabivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drabiv",
		TitleName: "德拉比夫",
		TitleCode: "b_drabiv",
	}
}
