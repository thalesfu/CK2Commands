package saravan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈什KhashBarony struct {
	feud.BaseBarony
}

var BKhash哈什 feud.Barony = &哈什KhashBarony{}

func init() {
	f := BKhash哈什.(*哈什KhashBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khash",
		TitleName: "哈什",
		TitleCode: "b_khash",
	}
}
