package nandana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库沙布KhushabBarony struct {
	feud.BaseBarony
}

var BKhushab库沙布 feud.Barony = &库沙布KhushabBarony{}

func init() {
    f := BKhushab库沙布.(*库沙布KhushabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khushab",
		TitleName: "库沙布",
		TitleCode: "b_khushab",
	}
}
