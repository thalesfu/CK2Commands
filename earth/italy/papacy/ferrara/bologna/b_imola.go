package bologna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊莫拉ImolaBarony struct {
	feud.BaseBarony
}

var BImola伊莫拉 feud.Barony = &伊莫拉ImolaBarony{}

func init() {
	f := BImola伊莫拉.(*伊莫拉ImolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "imola",
		TitleName: "伊莫拉",
		TitleCode: "b_imola",
	}
}
