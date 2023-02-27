package jersika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科克内塞KokneseBarony struct {
	feud.BaseBarony
}

var BKoknese科克内塞 feud.Barony = &科克内塞KokneseBarony{}

func init() {
    f := BKoknese科克内塞.(*科克内塞KokneseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koknese",
		TitleName: "科克内塞",
		TitleCode: "b_koknese",
	}
}
