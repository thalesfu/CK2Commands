package seletyteniz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科斯希KoschiBarony struct {
	feud.BaseBarony
}

var BKoschi科斯希 feud.Barony = &科斯希KoschiBarony{}

func init() {
	f := BKoschi科斯希.(*科斯希KoschiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koschi",
		TitleName: "科斯希",
		TitleCode: "b_koschi",
	}
}
