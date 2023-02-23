package cakrakuta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富度那揭罗PudunagramBarony struct {
	feud.BaseBarony
}

var BPudunagram富度那揭罗 feud.Barony = &富度那揭罗PudunagramBarony{}

func init() {
	f := BPudunagram富度那揭罗.(*富度那揭罗PudunagramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pudunagram",
		TitleName: "富度那揭罗",
		TitleCode: "b_pudunagram",
	}
}
