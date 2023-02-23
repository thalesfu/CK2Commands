package turgay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼卡尔MankalBarony struct {
	feud.BaseBarony
}

var BMankal曼卡尔 feud.Barony = &曼卡尔MankalBarony{}

func init() {
	f := BMankal曼卡尔.(*曼卡尔MankalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mankal",
		TitleName: "曼卡尔",
		TitleCode: "b_mankal",
	}
}
