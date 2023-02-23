package celle

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维廷根WittingenBarony struct {
	feud.BaseBarony
}

var BWittingen维廷根 feud.Barony = &维廷根WittingenBarony{}

func init() {
	f := BWittingen维廷根.(*维廷根WittingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wittingen",
		TitleName: "维廷根",
		TitleCode: "b_wittingen",
	}
}
