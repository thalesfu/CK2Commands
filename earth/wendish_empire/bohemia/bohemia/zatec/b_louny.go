package zatec

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛乌尼LounyBarony struct {
	feud.BaseBarony
}

var BLouny洛乌尼 feud.Barony = &洛乌尼LounyBarony{}

func init() {
    f := BLouny洛乌尼.(*洛乌尼LounyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "louny",
		TitleName: "洛乌尼",
		TitleCode: "b_louny",
	}
}
