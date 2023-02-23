package rummah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯苏叶雷AssuayerahBarony struct {
	feud.BaseBarony
}

var BAssuayerah阿斯苏叶雷 feud.Barony = &阿斯苏叶雷AssuayerahBarony{}

func init() {
	f := BAssuayerah阿斯苏叶雷.(*阿斯苏叶雷AssuayerahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "assuayerah",
		TitleName: "阿斯苏叶雷",
		TitleCode: "b_assuayerah",
	}
}
