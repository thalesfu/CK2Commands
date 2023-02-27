package piemonte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅萨拉诺MessaranoBarony struct {
	feud.BaseBarony
}

var BMessarano梅萨拉诺 feud.Barony = &梅萨拉诺MessaranoBarony{}

func init() {
    f := BMessarano梅萨拉诺.(*梅萨拉诺MessaranoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "messarano",
		TitleName: "梅萨拉诺",
		TitleCode: "b_messarano",
	}
}
