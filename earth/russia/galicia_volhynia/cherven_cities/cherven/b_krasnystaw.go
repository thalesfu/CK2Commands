package cherven

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉斯内斯塔夫KrasnystawBarony struct {
	feud.BaseBarony
}

var BKrasnystaw克拉斯内斯塔夫 feud.Barony = &克拉斯内斯塔夫KrasnystawBarony{}

func init() {
    f := BKrasnystaw克拉斯内斯塔夫.(*克拉斯内斯塔夫KrasnystawBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krasnystaw",
		TitleName: "克拉斯内斯塔夫",
		TitleCode: "b_krasnystaw",
	}
}
