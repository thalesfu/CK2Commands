package haruppeswara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡那兜KunnaturBarony struct {
	feud.BaseBarony
}

var BKunnatur贡那兜 feud.Barony = &贡那兜KunnaturBarony{}

func init() {
	f := BKunnatur贡那兜.(*贡那兜KunnaturBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kunnatur",
		TitleName: "贡那兜",
		TitleCode: "b_kunnatur",
	}
}
