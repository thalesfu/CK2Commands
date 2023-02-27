package nalsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内门奇内NemencineBarony struct {
	feud.BaseBarony
}

var BNemencine内门奇内 feud.Barony = &内门奇内NemencineBarony{}

func init() {
    f := BNemencine内门奇内.(*内门奇内NemencineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nemencine",
		TitleName: "内门奇内",
		TitleCode: "b_nemencine",
	}
}
