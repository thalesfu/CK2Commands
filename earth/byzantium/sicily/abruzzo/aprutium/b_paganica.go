package aprutium

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕格尼卡PaganicaBarony struct {
	feud.BaseBarony
}

var BPaganica帕格尼卡 feud.Barony = &帕格尼卡PaganicaBarony{}

func init() {
    f := BPaganica帕格尼卡.(*帕格尼卡PaganicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paganica",
		TitleName: "帕格尼卡",
		TitleCode: "b_paganica",
	}
}
