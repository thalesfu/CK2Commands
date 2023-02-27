package ural

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯季_卡塔夫UstkatavBarony struct {
	feud.BaseBarony
}

var BUstkatav乌斯季_卡塔夫 feud.Barony = &乌斯季_卡塔夫UstkatavBarony{}

func init() {
    f := BUstkatav乌斯季_卡塔夫.(*乌斯季_卡塔夫UstkatavBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ustkatav",
		TitleName: "乌斯季_卡塔夫",
		TitleCode: "b_ustkatav",
	}
}
