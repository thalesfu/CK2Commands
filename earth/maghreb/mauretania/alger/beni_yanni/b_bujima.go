package beni_yanni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布济马BujimaBarony struct {
	feud.BaseBarony
}

var BBujima布济马 feud.Barony = &布济马BujimaBarony{}

func init() {
    f := BBujima布济马.(*布济马BujimaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bujima",
		TitleName: "布济马",
		TitleCode: "b_bujima",
	}
}
