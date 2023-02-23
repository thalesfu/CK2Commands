package luristan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍拉马巴德KhorramabadBarony struct {
	feud.BaseBarony
}

var BKhorramabad霍拉马巴德 feud.Barony = &霍拉马巴德KhorramabadBarony{}

func init() {
	f := BKhorramabad霍拉马巴德.(*霍拉马巴德KhorramabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khorramabad",
		TitleName: "霍拉马巴德",
		TitleCode: "b_khorramabad",
	}
}
