package dihistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃先古雷GasankuliBarony struct {
	feud.BaseBarony
}

var BGasankuli埃先古雷 feud.Barony = &埃先古雷GasankuliBarony{}

func init() {
    f := BGasankuli埃先古雷.(*埃先古雷GasankuliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gasankuli",
		TitleName: "埃先古雷",
		TitleCode: "b_gasankuli",
	}
}
