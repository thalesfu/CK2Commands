package rugen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡伦察CharenzaBarony struct {
	feud.BaseBarony
}

var BCharenza卡伦察 feud.Barony = &卡伦察CharenzaBarony{}

func init() {
    f := BCharenza卡伦察.(*卡伦察CharenzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "charenza",
		TitleName: "卡伦察",
		TitleCode: "b_charenza",
	}
}
