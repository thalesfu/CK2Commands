package vijayapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗阇耶城VijayapuraBarony struct {
	feud.BaseBarony
}

var BVijayapura毗阇耶城 feud.Barony = &毗阇耶城VijayapuraBarony{}

func init() {
    f := BVijayapura毗阇耶城.(*毗阇耶城VijayapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vijayapura",
		TitleName: "毗阇耶城",
		TitleCode: "b_vijayapura",
	}
}
