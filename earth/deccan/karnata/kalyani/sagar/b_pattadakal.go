package sagar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕塔达卡尔PattadakalBarony struct {
	feud.BaseBarony
}

var BPattadakal帕塔达卡尔 feud.Barony = &帕塔达卡尔PattadakalBarony{}

func init() {
	f := BPattadakal帕塔达卡尔.(*帕塔达卡尔PattadakalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pattadakal",
		TitleName: "帕塔达卡尔",
		TitleCode: "b_pattadakal",
	}
}
