package calarasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔泰尼察OltenitaBarony struct {
	feud.BaseBarony
}

var BOltenita奥尔泰尼察 feud.Barony = &奥尔泰尼察OltenitaBarony{}

func init() {
	f := BOltenita奥尔泰尼察.(*奥尔泰尼察OltenitaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oltenita",
		TitleName: "奥尔泰尼察",
		TitleCode: "b_oltenita",
	}
}
