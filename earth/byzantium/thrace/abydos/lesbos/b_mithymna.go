package lesbos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米西姆纳MithymnaBarony struct {
	feud.BaseBarony
}

var BMithymna米西姆纳 feud.Barony = &米西姆纳MithymnaBarony{}

func init() {
    f := BMithymna米西姆纳.(*米西姆纳MithymnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mithymna",
		TitleName: "米西姆纳",
		TitleCode: "b_mithymna",
	}
}
