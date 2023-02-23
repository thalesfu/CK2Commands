package tashkurgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 行地HangdiBarony struct {
	feud.BaseBarony
}

var BHangdi行地 feud.Barony = &行地HangdiBarony{}

func init() {
	f := BHangdi行地.(*行地HangdiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hangdi",
		TitleName: "行地",
		TitleCode: "b_hangdi",
	}
}
