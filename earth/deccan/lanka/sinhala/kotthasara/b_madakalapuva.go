package kotthasara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩陀迦兰浮瓦MadakalapuvaBarony struct {
	feud.BaseBarony
}

var BMadakalapuva摩陀迦兰浮瓦 feud.Barony = &摩陀迦兰浮瓦MadakalapuvaBarony{}

func init() {
	f := BMadakalapuva摩陀迦兰浮瓦.(*摩陀迦兰浮瓦MadakalapuvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "madakalapuva",
		TitleName: "摩陀迦兰浮瓦",
		TitleCode: "b_madakalapuva",
	}
}
