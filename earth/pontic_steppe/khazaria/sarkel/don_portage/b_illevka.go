package don_portage

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊利耶夫卡IllevkaBarony struct {
	feud.BaseBarony
}

var BIllevka伊利耶夫卡 feud.Barony = &伊利耶夫卡IllevkaBarony{}

func init() {
    f := BIllevka伊利耶夫卡.(*伊利耶夫卡IllevkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "illevka",
		TitleName: "伊利耶夫卡",
		TitleCode: "b_illevka",
	}
}
