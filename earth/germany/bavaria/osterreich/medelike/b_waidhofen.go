package medelike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 魏德霍芬WaidhofenBarony struct {
	feud.BaseBarony
}

var BWaidhofen魏德霍芬 feud.Barony = &魏德霍芬WaidhofenBarony{}

func init() {
    f := BWaidhofen魏德霍芬.(*魏德霍芬WaidhofenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "waidhofen",
		TitleName: "魏德霍芬",
		TitleCode: "b_waidhofen",
	}
}
