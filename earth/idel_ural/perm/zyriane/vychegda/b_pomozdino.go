package vychegda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波莫兹季诺PomozdinoBarony struct {
	feud.BaseBarony
}

var BPomozdino波莫兹季诺 feud.Barony = &波莫兹季诺PomozdinoBarony{}

func init() {
    f := BPomozdino波莫兹季诺.(*波莫兹季诺PomozdinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pomozdino",
		TitleName: "波莫兹季诺",
		TitleCode: "b_pomozdino",
	}
}
