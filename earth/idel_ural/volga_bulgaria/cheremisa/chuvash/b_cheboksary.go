package chuvash

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切博克萨雷CheboksaryBarony struct {
	feud.BaseBarony
}

var BCheboksary切博克萨雷 feud.Barony = &切博克萨雷CheboksaryBarony{}

func init() {
    f := BCheboksary切博克萨雷.(*切博克萨雷CheboksaryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cheboksary",
		TitleName: "切博克萨雷",
		TitleCode: "b_cheboksary",
	}
}
