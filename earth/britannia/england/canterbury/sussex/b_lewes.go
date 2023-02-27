package sussex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 刘易斯LewesBarony struct {
	feud.BaseBarony
}

var BLewes刘易斯 feud.Barony = &刘易斯LewesBarony{}

func init() {
    f := BLewes刘易斯.(*刘易斯LewesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lewes",
		TitleName: "刘易斯",
		TitleCode: "b_lewes",
	}
}
