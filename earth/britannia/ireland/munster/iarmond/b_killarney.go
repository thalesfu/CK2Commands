package iarmond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基拉尼KillarneyBarony struct {
	feud.BaseBarony
}

var BKillarney基拉尼 feud.Barony = &基拉尼KillarneyBarony{}

func init() {
	f := BKillarney基拉尼.(*基拉尼KillarneyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "killarney",
		TitleName: "基拉尼",
		TitleCode: "b_killarney",
	}
}
