package cieszyn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别尔斯科BielskoBarony struct {
	feud.BaseBarony
}

var BBielsko别尔斯科 feud.Barony = &别尔斯科BielskoBarony{}

func init() {
    f := BBielsko别尔斯科.(*别尔斯科BielskoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bielsko",
		TitleName: "别尔斯科",
		TitleCode: "b_bielsko",
	}
}
