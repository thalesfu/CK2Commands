package djado

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图莫TummoBarony struct {
	feud.BaseBarony
}

var BTummo图莫 feud.Barony = &图莫TummoBarony{}

func init() {
	f := BTummo图莫.(*图莫TummoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tummo",
		TitleName: "图莫",
		TitleCode: "b_tummo",
	}
}
