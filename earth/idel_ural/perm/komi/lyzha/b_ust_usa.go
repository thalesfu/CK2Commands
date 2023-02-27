package lyzha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯季_乌萨Ust_usaBarony struct {
	feud.BaseBarony
}

var BUst_usa乌斯季_乌萨 feud.Barony = &乌斯季_乌萨Ust_usaBarony{}

func init() {
    f := BUst_usa乌斯季_乌萨.(*乌斯季_乌萨Ust_usaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ust_usa",
		TitleName: "乌斯季_乌萨",
		TitleCode: "b_ust_usa",
	}
}
