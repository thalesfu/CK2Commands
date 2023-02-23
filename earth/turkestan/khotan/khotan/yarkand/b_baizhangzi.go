package yarkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柏杖子BaizhangziBarony struct {
	feud.BaseBarony
}

var BBaizhangzi柏杖子 feud.Barony = &柏杖子BaizhangziBarony{}

func init() {
	f := BBaizhangzi柏杖子.(*柏杖子BaizhangziBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baizhangzi",
		TitleName: "柏杖子",
		TitleCode: "b_baizhangzi",
	}
}
