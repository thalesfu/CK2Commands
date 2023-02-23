package medantaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑耆SanjooBarony struct {
	feud.BaseBarony
}

var BSanjoo桑耆 feud.Barony = &桑耆SanjooBarony{}

func init() {
	f := BSanjoo桑耆.(*桑耆SanjooBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanjoo",
		TitleName: "桑耆",
		TitleCode: "b_sanjoo",
	}
}
