package cherchen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏塘SutangBarony struct {
	feud.BaseBarony
}

var BSutang苏塘 feud.Barony = &苏塘SutangBarony{}

func init() {
    f := BSutang苏塘.(*苏塘SutangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sutang",
		TitleName: "苏塘",
		TitleCode: "b_sutang",
	}
}
