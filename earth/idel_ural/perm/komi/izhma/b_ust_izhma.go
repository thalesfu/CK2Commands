package izhma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯季_伊日马Ust_izhmaBarony struct {
	feud.BaseBarony
}

var BUst_izhma乌斯季_伊日马 feud.Barony = &乌斯季_伊日马Ust_izhmaBarony{}

func init() {
    f := BUst_izhma乌斯季_伊日马.(*乌斯季_伊日马Ust_izhmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ust_izhma",
		TitleName: "乌斯季_伊日马",
		TitleCode: "b_ust_izhma",
	}
}
