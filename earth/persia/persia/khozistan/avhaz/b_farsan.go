package avhaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法尔桑FarsanBarony struct {
	feud.BaseBarony
}

var BFarsan法尔桑 feud.Barony = &法尔桑FarsanBarony{}

func init() {
    f := BFarsan法尔桑.(*法尔桑FarsanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "farsan",
		TitleName: "法尔桑",
		TitleCode: "b_farsan",
	}
}
