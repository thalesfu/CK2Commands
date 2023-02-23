package arta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多莫卡斯特朗ThomokastronBarony struct {
	feud.BaseBarony
}

var BThomokastron多莫卡斯特朗 feud.Barony = &多莫卡斯特朗ThomokastronBarony{}

func init() {
	f := BThomokastron多莫卡斯特朗.(*多莫卡斯特朗ThomokastronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thomokastron",
		TitleName: "多莫卡斯特朗",
		TitleCode: "b_thomokastron",
	}
}
