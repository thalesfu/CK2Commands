package boulogne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣波勒SaintpolBarony struct {
	feud.BaseBarony
}

var BSaintpol圣波勒 feud.Barony = &圣波勒SaintpolBarony{}

func init() {
	f := BSaintpol圣波勒.(*圣波勒SaintpolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saintpol",
		TitleName: "圣波勒",
		TitleCode: "b_saintpol",
	}
}
