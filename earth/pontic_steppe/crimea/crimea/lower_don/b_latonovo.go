package lower_don

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉托诺沃LatonovoBarony struct {
	feud.BaseBarony
}

var BLatonovo拉托诺沃 feud.Barony = &拉托诺沃LatonovoBarony{}

func init() {
    f := BLatonovo拉托诺沃.(*拉托诺沃LatonovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "latonovo",
		TitleName: "拉托诺沃",
		TitleCode: "b_latonovo",
	}
}
