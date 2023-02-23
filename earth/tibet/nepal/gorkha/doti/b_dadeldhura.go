package doti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达代尔图拉DadeldhuraBarony struct {
	feud.BaseBarony
}

var BDadeldhura达代尔图拉 feud.Barony = &达代尔图拉DadeldhuraBarony{}

func init() {
	f := BDadeldhura达代尔图拉.(*达代尔图拉DadeldhuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dadeldhura",
		TitleName: "达代尔图拉",
		TitleCode: "b_dadeldhura",
	}
}
