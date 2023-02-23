package livs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利胡拉LihulaBarony struct {
	feud.BaseBarony
}

var BLihula利胡拉 feud.Barony = &利胡拉LihulaBarony{}

func init() {
	f := BLihula利胡拉.(*利胡拉LihulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lihula",
		TitleName: "利胡拉",
		TitleCode: "b_lihula",
	}
}
