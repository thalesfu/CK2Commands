package gyantse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 白居PalchoBarony struct {
	feud.BaseBarony
}

var BPalcho白居 feud.Barony = &白居PalchoBarony{}

func init() {
    f := BPalcho白居.(*白居PalchoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palcho",
		TitleName: "白居",
		TitleCode: "b_palcho",
	}
}
