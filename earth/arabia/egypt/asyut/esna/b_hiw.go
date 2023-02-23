package esna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希乌HiwBarony struct {
	feud.BaseBarony
}

var BHiw希乌 feud.Barony = &希乌HiwBarony{}

func init() {
	f := BHiw希乌.(*希乌HiwBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hiw",
		TitleName: "希乌",
		TitleCode: "b_hiw",
	}
}
