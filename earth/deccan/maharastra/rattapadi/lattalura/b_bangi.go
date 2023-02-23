package lattalura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 槃耆BangiBarony struct {
	feud.BaseBarony
}

var BBangi槃耆 feud.Barony = &槃耆BangiBarony{}

func init() {
	f := BBangi槃耆.(*槃耆BangiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bangi",
		TitleName: "槃耆",
		TitleCode: "b_bangi",
	}
}
