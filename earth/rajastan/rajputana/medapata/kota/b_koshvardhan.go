package kota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 俱舍伐弹那KoshvardhanBarony struct {
	feud.BaseBarony
}

var BKoshvardhan俱舍伐弹那 feud.Barony = &俱舍伐弹那KoshvardhanBarony{}

func init() {
    f := BKoshvardhan俱舍伐弹那.(*俱舍伐弹那KoshvardhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koshvardhan",
		TitleName: "俱舍伐弹那",
		TitleCode: "b_koshvardhan",
	}
}
