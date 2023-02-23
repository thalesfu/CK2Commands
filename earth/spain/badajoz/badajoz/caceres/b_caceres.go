package caceres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡塞雷斯CaceresBarony struct {
	feud.BaseBarony
}

var BCaceres卡塞雷斯 feud.Barony = &卡塞雷斯CaceresBarony{}

func init() {
	f := BCaceres卡塞雷斯.(*卡塞雷斯CaceresBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caceres",
		TitleName: "卡塞雷斯",
		TitleCode: "b_caceres",
	}
}
