package bulgar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏瓦尔SuarBarony struct {
	feud.BaseBarony
}

var BSuar苏瓦尔 feud.Barony = &苏瓦尔SuarBarony{}

func init() {
    f := BSuar苏瓦尔.(*苏瓦尔SuarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suar",
		TitleName: "苏瓦尔",
		TitleCode: "b_suar",
	}
}
