package soli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比利诺波耶Bilino_poljeBarony struct {
	feud.BaseBarony
}

var BBilino_polje比利诺波耶 feud.Barony = &比利诺波耶Bilino_poljeBarony{}

func init() {
    f := BBilino_polje比利诺波耶.(*比利诺波耶Bilino_poljeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bilino_polje",
		TitleName: "比利诺波耶",
		TitleCode: "b_bilino_polje",
	}
}
