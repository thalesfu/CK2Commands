package karmanta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢艺婆旃LatembarcemBarony struct {
	feud.BaseBarony
}

var BLatembarcem卢艺婆旃 feud.Barony = &卢艺婆旃LatembarcemBarony{}

func init() {
	f := BLatembarcem卢艺婆旃.(*卢艺婆旃LatembarcemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "latembarcem",
		TitleName: "卢艺婆旃",
		TitleCode: "b_latembarcem",
	}
}
