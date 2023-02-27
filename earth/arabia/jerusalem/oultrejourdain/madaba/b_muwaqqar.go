package madaba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆瓦盖尔MuwaqqarBarony struct {
	feud.BaseBarony
}

var BMuwaqqar穆瓦盖尔 feud.Barony = &穆瓦盖尔MuwaqqarBarony{}

func init() {
    f := BMuwaqqar穆瓦盖尔.(*穆瓦盖尔MuwaqqarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muwaqqar",
		TitleName: "穆瓦盖尔",
		TitleCode: "b_muwaqqar",
	}
}
