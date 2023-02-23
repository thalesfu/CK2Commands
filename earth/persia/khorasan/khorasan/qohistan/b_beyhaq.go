package qohistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜哈格BeyhaqBarony struct {
	feud.BaseBarony
}

var BBeyhaq拜哈格 feud.Barony = &拜哈格BeyhaqBarony{}

func init() {
	f := BBeyhaq拜哈格.(*拜哈格BeyhaqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beyhaq",
		TitleName: "拜哈格",
		TitleCode: "b_beyhaq",
	}
}
