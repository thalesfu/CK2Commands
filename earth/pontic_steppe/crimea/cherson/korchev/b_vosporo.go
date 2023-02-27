package korchev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃斯波罗VosporoBarony struct {
	feud.BaseBarony
}

var BVosporo沃斯波罗 feud.Barony = &沃斯波罗VosporoBarony{}

func init() {
    f := BVosporo沃斯波罗.(*沃斯波罗VosporoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vosporo",
		TitleName: "沃斯波罗",
		TitleCode: "b_vosporo",
	}
}
