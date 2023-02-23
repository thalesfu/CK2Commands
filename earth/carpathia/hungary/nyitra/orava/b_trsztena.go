package orava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特尔斯泰纳TrsztenaBarony struct {
	feud.BaseBarony
}

var BTrsztena特尔斯泰纳 feud.Barony = &特尔斯泰纳TrsztenaBarony{}

func init() {
	f := BTrsztena特尔斯泰纳.(*特尔斯泰纳TrsztenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trsztena",
		TitleName: "特尔斯泰纳",
		TitleCode: "b_trsztena",
	}
}
