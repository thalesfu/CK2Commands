package lolland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马里博MariboBarony struct {
	feud.BaseBarony
}

var BMaribo马里博 feud.Barony = &马里博MariboBarony{}

func init() {
	f := BMaribo马里博.(*马里博MariboBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maribo",
		TitleName: "马里博",
		TitleCode: "b_maribo",
	}
}
