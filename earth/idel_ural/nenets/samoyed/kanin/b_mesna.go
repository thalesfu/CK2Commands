package kanin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅斯纳MesnaBarony struct {
	feud.BaseBarony
}

var BMesna梅斯纳 feud.Barony = &梅斯纳MesnaBarony{}

func init() {
    f := BMesna梅斯纳.(*梅斯纳MesnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mesna",
		TitleName: "梅斯纳",
		TitleCode: "b_mesna",
	}
}
