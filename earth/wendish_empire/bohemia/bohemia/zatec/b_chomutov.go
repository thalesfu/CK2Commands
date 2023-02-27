package zatec

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍穆托夫ChomutovBarony struct {
	feud.BaseBarony
}

var BChomutov霍穆托夫 feud.Barony = &霍穆托夫ChomutovBarony{}

func init() {
    f := BChomutov霍穆托夫.(*霍穆托夫ChomutovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chomutov",
		TitleName: "霍穆托夫",
		TitleCode: "b_chomutov",
	}
}
