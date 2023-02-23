package arta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安格洛卡斯特罗AngelokastronBarony struct {
	feud.BaseBarony
}

var BAngelokastron安格洛卡斯特罗 feud.Barony = &安格洛卡斯特罗AngelokastronBarony{}

func init() {
	f := BAngelokastron安格洛卡斯特罗.(*安格洛卡斯特罗AngelokastronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "angelokastron",
		TitleName: "安格洛卡斯特罗",
		TitleCode: "b_angelokastron",
	}
}
