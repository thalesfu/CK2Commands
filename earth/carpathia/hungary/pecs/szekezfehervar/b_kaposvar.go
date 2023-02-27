package szekezfehervar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 考波什堡KaposvarBarony struct {
	feud.BaseBarony
}

var BKaposvar考波什堡 feud.Barony = &考波什堡KaposvarBarony{}

func init() {
    f := BKaposvar考波什堡.(*考波什堡KaposvarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaposvar",
		TitleName: "考波什堡",
		TitleCode: "b_kaposvar",
	}
}
