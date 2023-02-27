package mountain_cheremisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯塔尼奇纳亚StanichnayaBarony struct {
	feud.BaseBarony
}

var BStanichnaya斯塔尼奇纳亚 feud.Barony = &斯塔尼奇纳亚StanichnayaBarony{}

func init() {
    f := BStanichnaya斯塔尼奇纳亚.(*斯塔尼奇纳亚StanichnayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stanichnaya",
		TitleName: "斯塔尼奇纳亚",
		TitleCode: "b_stanichnaya",
	}
}
