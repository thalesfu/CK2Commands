package uzens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈桑KhasanBarony struct {
	feud.BaseBarony
}

var BKhasan哈桑 feud.Barony = &哈桑KhasanBarony{}

func init() {
    f := BKhasan哈桑.(*哈桑KhasanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khasan",
		TitleName: "哈桑",
		TitleCode: "b_khasan",
	}
}
