package loon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法尔肯堡ValkenburgBarony struct {
	feud.BaseBarony
}

var BValkenburg法尔肯堡 feud.Barony = &法尔肯堡ValkenburgBarony{}

func init() {
    f := BValkenburg法尔肯堡.(*法尔肯堡ValkenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valkenburg",
		TitleName: "法尔肯堡",
		TitleCode: "b_valkenburg",
	}
}
