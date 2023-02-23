package yazd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼尔NirBarony struct {
	feud.BaseBarony
}

var BNir尼尔 feud.Barony = &尼尔NirBarony{}

func init() {
	f := BNir尼尔.(*尼尔NirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nir",
		TitleName: "尼尔",
		TitleCode: "b_nir",
	}
}
