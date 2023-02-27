package kartaly

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎列奇内ZarechnyBarony struct {
	feud.BaseBarony
}

var BZarechny扎列奇内 feud.Barony = &扎列奇内ZarechnyBarony{}

func init() {
    f := BZarechny扎列奇内.(*扎列奇内ZarechnyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zarechny",
		TitleName: "扎列奇内",
		TitleCode: "b_zarechny",
	}
}
