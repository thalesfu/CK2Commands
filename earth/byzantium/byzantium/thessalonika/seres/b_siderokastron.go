package seres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡季罗卡斯特龙SiderokastronBarony struct {
	feud.BaseBarony
}

var BSiderokastron锡季罗卡斯特龙 feud.Barony = &锡季罗卡斯特龙SiderokastronBarony{}

func init() {
    f := BSiderokastron锡季罗卡斯特龙.(*锡季罗卡斯特龙SiderokastronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siderokastron",
		TitleName: "锡季罗卡斯特龙",
		TitleCode: "b_siderokastron",
	}
}
