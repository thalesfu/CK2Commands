package syrte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑙费利耶NawfalliyahBarony struct {
	feud.BaseBarony
}

var BNawfalliyah瑙费利耶 feud.Barony = &瑙费利耶NawfalliyahBarony{}

func init() {
    f := BNawfalliyah瑙费利耶.(*瑙费利耶NawfalliyahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nawfalliyah",
		TitleName: "瑙费利耶",
		TitleCode: "b_nawfalliyah",
	}
}
