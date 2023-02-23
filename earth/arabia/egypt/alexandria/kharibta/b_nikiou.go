package kharibta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼基乌NikiouBarony struct {
	feud.BaseBarony
}

var BNikiou尼基乌 feud.Barony = &尼基乌NikiouBarony{}

func init() {
	f := BNikiou尼基乌.(*尼基乌NikiouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nikiou",
		TitleName: "尼基乌",
		TitleCode: "b_nikiou",
	}
}
