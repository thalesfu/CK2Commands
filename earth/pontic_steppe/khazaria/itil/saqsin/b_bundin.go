package saqsin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 本丁BundinBarony struct {
	feud.BaseBarony
}

var BBundin本丁 feud.Barony = &本丁BundinBarony{}

func init() {
    f := BBundin本丁.(*本丁BundinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bundin",
		TitleName: "本丁",
		TitleCode: "b_bundin",
	}
}
