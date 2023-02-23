package aleppo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈阿赖努阿曼MaaratannumanBarony struct {
	feud.BaseBarony
}

var BMaaratannuman迈阿赖努阿曼 feud.Barony = &迈阿赖努阿曼MaaratannumanBarony{}

func init() {
	f := BMaaratannuman迈阿赖努阿曼.(*迈阿赖努阿曼MaaratannumanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maaratannuman",
		TitleName: "迈阿赖努阿曼",
		TitleCode: "b_maaratannuman",
	}
}
