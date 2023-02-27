package esfahan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 汉萨尔KhansarBarony struct {
	feud.BaseBarony
}

var BKhansar汉萨尔 feud.Barony = &汉萨尔KhansarBarony{}

func init() {
    f := BKhansar汉萨尔.(*汉萨尔KhansarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khansar",
		TitleName: "汉萨尔",
		TitleCode: "b_khansar",
	}
}
