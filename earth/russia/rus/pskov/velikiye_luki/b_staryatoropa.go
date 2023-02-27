package velikiye_luki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旧托罗帕StaryatoropaBarony struct {
	feud.BaseBarony
}

var BStaryatoropa旧托罗帕 feud.Barony = &旧托罗帕StaryatoropaBarony{}

func init() {
    f := BStaryatoropa旧托罗帕.(*旧托罗帕StaryatoropaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "staryatoropa",
		TitleName: "旧托罗帕",
		TitleCode: "b_staryatoropa",
	}
}
