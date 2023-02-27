package saintois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索尔西圣马尔坦SorcystmartinBarony struct {
	feud.BaseBarony
}

var BSorcystmartin索尔西圣马尔坦 feud.Barony = &索尔西圣马尔坦SorcystmartinBarony{}

func init() {
    f := BSorcystmartin索尔西圣马尔坦.(*索尔西圣马尔坦SorcystmartinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sorcystmartin",
		TitleName: "索尔西圣马尔坦",
		TitleCode: "b_sorcystmartin",
	}
}
