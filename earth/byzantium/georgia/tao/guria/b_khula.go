package guria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡拉KhulaBarony struct {
	feud.BaseBarony
}

var BKhula胡拉 feud.Barony = &胡拉KhulaBarony{}

func init() {
	f := BKhula胡拉.(*胡拉KhulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khula",
		TitleName: "胡拉",
		TitleCode: "b_khula",
	}
}
