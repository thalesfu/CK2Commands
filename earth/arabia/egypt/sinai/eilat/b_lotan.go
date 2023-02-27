package eilat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗坍LotanBarony struct {
	feud.BaseBarony
}

var BLotan罗坍 feud.Barony = &罗坍LotanBarony{}

func init() {
    f := BLotan罗坍.(*罗坍LotanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lotan",
		TitleName: "罗坍",
		TitleCode: "b_lotan",
	}
}
