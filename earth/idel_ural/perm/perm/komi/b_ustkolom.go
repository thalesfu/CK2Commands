package komi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯季_科洛姆UstkolomBarony struct {
	feud.BaseBarony
}

var BUstkolom乌斯季_科洛姆 feud.Barony = &乌斯季_科洛姆UstkolomBarony{}

func init() {
    f := BUstkolom乌斯季_科洛姆.(*乌斯季_科洛姆UstkolomBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ustkolom",
		TitleName: "乌斯季_科洛姆",
		TitleCode: "b_ustkolom",
	}
}
