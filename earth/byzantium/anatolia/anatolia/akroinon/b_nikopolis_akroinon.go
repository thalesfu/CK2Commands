package akroinon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼科波利斯Nikopolis_akroinonBarony struct {
	feud.BaseBarony
}

var BNikopolis_akroinon尼科波利斯 feud.Barony = &尼科波利斯Nikopolis_akroinonBarony{}

func init() {
    f := BNikopolis_akroinon尼科波利斯.(*尼科波利斯Nikopolis_akroinonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nikopolis_akroinon",
		TitleName: "尼科波利斯",
		TitleCode: "b_nikopolis_akroinon",
	}
}
