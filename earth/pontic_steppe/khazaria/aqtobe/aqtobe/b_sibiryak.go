package aqtobe

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 失必儿雅克SibiryakBarony struct {
	feud.BaseBarony
}

var BSibiryak失必儿雅克 feud.Barony = &失必儿雅克SibiryakBarony{}

func init() {
    f := BSibiryak失必儿雅克.(*失必儿雅克SibiryakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sibiryak",
		TitleName: "失必儿雅克",
		TitleCode: "b_sibiryak",
	}
}
