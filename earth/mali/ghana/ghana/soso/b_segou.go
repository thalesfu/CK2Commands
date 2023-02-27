package soso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞古SegouBarony struct {
	feud.BaseBarony
}

var BSegou塞古 feud.Barony = &塞古SegouBarony{}

func init() {
    f := BSegou塞古.(*塞古SegouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "segou",
		TitleName: "塞古",
		TitleCode: "b_segou",
	}
}
