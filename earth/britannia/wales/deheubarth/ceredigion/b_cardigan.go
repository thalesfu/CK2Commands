package ceredigion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡迪根CardiganBarony struct {
	feud.BaseBarony
}

var BCardigan卡迪根 feud.Barony = &卡迪根CardiganBarony{}

func init() {
    f := BCardigan卡迪根.(*卡迪根CardiganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cardigan",
		TitleName: "卡迪根",
		TitleCode: "b_cardigan",
	}
}
