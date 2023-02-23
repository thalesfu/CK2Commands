package pitten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮滕PittenBarony struct {
	feud.BaseBarony
}

var BPitten皮滕 feud.Barony = &皮滕PittenBarony{}

func init() {
	f := BPitten皮滕.(*皮滕PittenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pitten",
		TitleName: "皮滕",
		TitleCode: "b_pitten",
	}
}
