package qom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 豪尔萨巴德KhourabadBarony struct {
	feud.BaseBarony
}

var BKhourabad豪尔萨巴德 feud.Barony = &豪尔萨巴德KhourabadBarony{}

func init() {
	f := BKhourabad豪尔萨巴德.(*豪尔萨巴德KhourabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khourabad",
		TitleName: "豪尔萨巴德",
		TitleCode: "b_khourabad",
	}
}
