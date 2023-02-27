package vel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列姆特博日LemtybozhBarony struct {
	feud.BaseBarony
}

var BLemtybozh列姆特博日 feud.Barony = &列姆特博日LemtybozhBarony{}

func init() {
    f := BLemtybozh列姆特博日.(*列姆特博日LemtybozhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lemtybozh",
		TitleName: "列姆特博日",
		TitleCode: "b_lemtybozh",
	}
}
