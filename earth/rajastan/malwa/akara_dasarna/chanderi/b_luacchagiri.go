package chanderi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 捋车耆厘LuacchagiriBarony struct {
	feud.BaseBarony
}

var BLuacchagiri捋车耆厘 feud.Barony = &捋车耆厘LuacchagiriBarony{}

func init() {
    f := BLuacchagiri捋车耆厘.(*捋车耆厘LuacchagiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luacchagiri",
		TitleName: "捋车耆厘",
		TitleCode: "b_luacchagiri",
	}
}
