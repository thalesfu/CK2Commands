package ghana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昆比萨利赫Koumbi_salehBarony struct {
	feud.BaseBarony
}

var BKoumbi_saleh昆比萨利赫 feud.Barony = &昆比萨利赫Koumbi_salehBarony{}

func init() {
    f := BKoumbi_saleh昆比萨利赫.(*昆比萨利赫Koumbi_salehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koumbi_saleh",
		TitleName: "昆比萨利赫",
		TitleCode: "b_koumbi_saleh",
	}
}
