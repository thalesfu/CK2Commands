package turgay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯季_图尔盖Ust_turgayBarony struct {
	feud.BaseBarony
}

var BUst_turgay乌斯季_图尔盖 feud.Barony = &乌斯季_图尔盖Ust_turgayBarony{}

func init() {
    f := BUst_turgay乌斯季_图尔盖.(*乌斯季_图尔盖Ust_turgayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ust_turgay",
		TitleName: "乌斯季_图尔盖",
		TitleCode: "b_ust_turgay",
	}
}
