package tsagaannuur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈里温KhaliunBarony struct {
	feud.BaseBarony
}

var BKhaliun哈里温 feud.Barony = &哈里温KhaliunBarony{}

func init() {
    f := BKhaliun哈里温.(*哈里温KhaliunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khaliun",
		TitleName: "哈里温",
		TitleCode: "b_khaliun",
	}
}
