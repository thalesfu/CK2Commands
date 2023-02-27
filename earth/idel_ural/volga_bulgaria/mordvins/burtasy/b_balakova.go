package burtasy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴拉科沃BalakovaBarony struct {
	feud.BaseBarony
}

var BBalakova巴拉科沃 feud.Barony = &巴拉科沃BalakovaBarony{}

func init() {
    f := BBalakova巴拉科沃.(*巴拉科沃BalakovaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balakova",
		TitleName: "巴拉科沃",
		TitleCode: "b_balakova",
	}
}
