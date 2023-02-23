package zaragoza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅迪亚纳德亚拉贡MedianadearagonBarony struct {
	feud.BaseBarony
}

var BMedianadearagon梅迪亚纳德亚拉贡 feud.Barony = &梅迪亚纳德亚拉贡MedianadearagonBarony{}

func init() {
	f := BMedianadearagon梅迪亚纳德亚拉贡.(*梅迪亚纳德亚拉贡MedianadearagonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medianadearagon",
		TitleName: "梅迪亚纳德亚拉贡",
		TitleCode: "b_medianadearagon",
	}
}
