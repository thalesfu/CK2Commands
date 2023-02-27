package tadmekka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯卡科EskakoBarony struct {
	feud.BaseBarony
}

var BEskako埃斯卡科 feud.Barony = &埃斯卡科EskakoBarony{}

func init() {
    f := BEskako埃斯卡科.(*埃斯卡科EskakoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eskako",
		TitleName: "埃斯卡科",
		TitleCode: "b_eskako",
	}
}
