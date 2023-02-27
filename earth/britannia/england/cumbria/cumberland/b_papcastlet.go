package cumberland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕普卡斯特尔PapcastletBarony struct {
	feud.BaseBarony
}

var BPapcastlet帕普卡斯特尔 feud.Barony = &帕普卡斯特尔PapcastletBarony{}

func init() {
    f := BPapcastlet帕普卡斯特尔.(*帕普卡斯特尔PapcastletBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "papcastlet",
		TitleName: "帕普卡斯特尔",
		TitleCode: "b_papcastlet",
	}
}
