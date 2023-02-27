package asturias_de_oviedo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希洪GijonBarony struct {
	feud.BaseBarony
}

var BGijon希洪 feud.Barony = &希洪GijonBarony{}

func init() {
    f := BGijon希洪.(*希洪GijonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gijon",
		TitleName: "希洪",
		TitleCode: "b_gijon",
	}
}
