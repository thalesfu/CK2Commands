package constantia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅吉迪亚MesgidiaBarony struct {
	feud.BaseBarony
}

var BMesgidia梅吉迪亚 feud.Barony = &梅吉迪亚MesgidiaBarony{}

func init() {
    f := BMesgidia梅吉迪亚.(*梅吉迪亚MesgidiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mesgidia",
		TitleName: "梅吉迪亚",
		TitleCode: "b_mesgidia",
	}
}
