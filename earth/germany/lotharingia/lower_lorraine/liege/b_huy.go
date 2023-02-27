package liege

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 于伊HuyBarony struct {
	feud.BaseBarony
}

var BHuy于伊 feud.Barony = &于伊HuyBarony{}

func init() {
    f := BHuy于伊.(*于伊HuyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "huy",
		TitleName: "于伊",
		TitleCode: "b_huy",
	}
}
