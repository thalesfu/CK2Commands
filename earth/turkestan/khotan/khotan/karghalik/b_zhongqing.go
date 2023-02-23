package karghalik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 仲庆ZhongqingBarony struct {
	feud.BaseBarony
}

var BZhongqing仲庆 feud.Barony = &仲庆ZhongqingBarony{}

func init() {
	f := BZhongqing仲庆.(*仲庆ZhongqingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhongqing",
		TitleName: "仲庆",
		TitleCode: "b_zhongqing",
	}
}
