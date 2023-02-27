package multan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玛哈图木拉希德Makhdum_rashidBarony struct {
	feud.BaseBarony
}

var BMakhdum_rashid玛哈图木拉希德 feud.Barony = &玛哈图木拉希德Makhdum_rashidBarony{}

func init() {
    f := BMakhdum_rashid玛哈图木拉希德.(*玛哈图木拉希德Makhdum_rashidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "makhdum_rashid",
		TitleName: "玛哈图木拉希德",
		TitleCode: "b_makhdum_rashid",
	}
}
