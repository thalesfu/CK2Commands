package ugra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃伊沃日VoyvozhBarony struct {
	feud.BaseBarony
}

var BVoyvozh沃伊沃日 feud.Barony = &沃伊沃日VoyvozhBarony{}

func init() {
    f := BVoyvozh沃伊沃日.(*沃伊沃日VoyvozhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "voyvozh",
		TitleName: "沃伊沃日",
		TitleCode: "b_voyvozh",
	}
}
