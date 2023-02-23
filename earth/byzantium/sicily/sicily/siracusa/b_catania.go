package siracusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡塔尼亚CataniaBarony struct {
	feud.BaseBarony
}

var BCatania卡塔尼亚 feud.Barony = &卡塔尼亚CataniaBarony{}

func init() {
	f := BCatania卡塔尼亚.(*卡塔尼亚CataniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "catania",
		TitleName: "卡塔尼亚",
		TitleCode: "b_catania",
	}
}
