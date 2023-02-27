package suenik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿雷尼AreniBarony struct {
	feud.BaseBarony
}

var BAreni阿雷尼 feud.Barony = &阿雷尼AreniBarony{}

func init() {
    f := BAreni阿雷尼.(*阿雷尼AreniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "areni",
		TitleName: "阿雷尼",
		TitleCode: "b_areni",
	}
}
