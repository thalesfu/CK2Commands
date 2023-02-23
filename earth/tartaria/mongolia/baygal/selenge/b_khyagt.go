package selenge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰克图KhyagtBarony struct {
	feud.BaseBarony
}

var BKhyagt恰克图 feud.Barony = &恰克图KhyagtBarony{}

func init() {
	f := BKhyagt恰克图.(*恰克图KhyagtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khyagt",
		TitleName: "恰克图",
		TitleCode: "b_khyagt",
	}
}
