package mtskheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡奥尼SioniBarony struct {
	feud.BaseBarony
}

var BSioni锡奥尼 feud.Barony = &锡奥尼SioniBarony{}

func init() {
	f := BSioni锡奥尼.(*锡奥尼SioniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sioni",
		TitleName: "锡奥尼",
		TitleCode: "b_sioni",
	}
}
