package kharibta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾特菲赫AtfihBarony struct {
	feud.BaseBarony
}

var BAtfih艾特菲赫 feud.Barony = &艾特菲赫AtfihBarony{}

func init() {
	f := BAtfih艾特菲赫.(*艾特菲赫AtfihBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "atfih",
		TitleName: "艾特菲赫",
		TitleCode: "b_atfih",
	}
}
