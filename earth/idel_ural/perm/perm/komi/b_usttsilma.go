package komi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯季_齐利马UsttsilmaBarony struct {
	feud.BaseBarony
}

var BUsttsilma乌斯季_齐利马 feud.Barony = &乌斯季_齐利马UsttsilmaBarony{}

func init() {
    f := BUsttsilma乌斯季_齐利马.(*乌斯季_齐利马UsttsilmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "usttsilma",
		TitleName: "乌斯季_齐利马",
		TitleCode: "b_usttsilma",
	}
}
