package lyncestis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普里莱普PrilapBarony struct {
	feud.BaseBarony
}

var BPrilap普里莱普 feud.Barony = &普里莱普PrilapBarony{}

func init() {
	f := BPrilap普里莱普.(*普里莱普PrilapBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prilap",
		TitleName: "普里莱普",
		TitleCode: "b_prilap",
	}
}
