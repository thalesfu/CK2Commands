package oldenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒宁根LoningenBarony struct {
	feud.BaseBarony
}

var BLoningen勒宁根 feud.Barony = &勒宁根LoningenBarony{}

func init() {
	f := BLoningen勒宁根.(*勒宁根LoningenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loningen",
		TitleName: "勒宁根",
		TitleCode: "b_loningen",
	}
}
