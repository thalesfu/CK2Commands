package rajanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗旬补罗RajanpurBarony struct {
	feud.BaseBarony
}

var BRajanpur罗旬补罗 feud.Barony = &罗旬补罗RajanpurBarony{}

func init() {
	f := BRajanpur罗旬补罗.(*罗旬补罗RajanpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rajanpur",
		TitleName: "罗旬补罗",
		TitleCode: "b_rajanpur",
	}
}
