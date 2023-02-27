package tana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡多罗夫卡GundorovkaBarony struct {
	feud.BaseBarony
}

var BGundorovka贡多罗夫卡 feud.Barony = &贡多罗夫卡GundorovkaBarony{}

func init() {
    f := BGundorovka贡多罗夫卡.(*贡多罗夫卡GundorovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gundorovka",
		TitleName: "贡多罗夫卡",
		TitleCode: "b_gundorovka",
	}
}
