package sarqihya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宰加济格ZagazigBarony struct {
	feud.BaseBarony
}

var BZagazig宰加济格 feud.Barony = &宰加济格ZagazigBarony{}

func init() {
	f := BZagazig宰加济格.(*宰加济格ZagazigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zagazig",
		TitleName: "宰加济格",
		TitleCode: "b_zagazig",
	}
}
