package mandapika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔达曼BordhomanBarony struct {
	feud.BaseBarony
}

var BBordhoman巴尔达曼 feud.Barony = &巴尔达曼BordhomanBarony{}

func init() {
	f := BBordhoman巴尔达曼.(*巴尔达曼BordhomanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bordhoman",
		TitleName: "巴尔达曼",
		TitleCode: "b_bordhoman",
	}
}
