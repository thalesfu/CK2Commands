package bhakkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉卡纳LarkanaBarony struct {
	feud.BaseBarony
}

var BLarkana拉卡纳 feud.Barony = &拉卡纳LarkanaBarony{}

func init() {
	f := BLarkana拉卡纳.(*拉卡纳LarkanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "larkana",
		TitleName: "拉卡纳",
		TitleCode: "b_larkana",
	}
}
