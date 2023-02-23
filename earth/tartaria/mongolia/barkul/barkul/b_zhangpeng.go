package barkul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 漳澎ZhangpengBarony struct {
	feud.BaseBarony
}

var BZhangpeng漳澎 feud.Barony = &漳澎ZhangpengBarony{}

func init() {
	f := BZhangpeng漳澎.(*漳澎ZhangpengBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhangpeng",
		TitleName: "漳澎",
		TitleCode: "b_zhangpeng",
	}
}
