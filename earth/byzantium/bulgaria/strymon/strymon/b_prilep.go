package strymon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普里莱普PrilepBarony struct {
	feud.BaseBarony
}

var BPrilep普里莱普 feud.Barony = &普里莱普PrilepBarony{}

func init() {
	f := BPrilep普里莱普.(*普里莱普PrilepBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prilep",
		TitleName: "普里莱普",
		TitleCode: "b_prilep",
	}
}
