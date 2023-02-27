package pinsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢尼涅茨LuninetsBarony struct {
	feud.BaseBarony
}

var BLuninets卢尼涅茨 feud.Barony = &卢尼涅茨LuninetsBarony{}

func init() {
    f := BLuninets卢尼涅茨.(*卢尼涅茨LuninetsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luninets",
		TitleName: "卢尼涅茨",
		TitleCode: "b_luninets",
	}
}
