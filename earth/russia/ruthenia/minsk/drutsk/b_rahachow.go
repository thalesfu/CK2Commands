package drutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗加乔夫RahachowBarony struct {
	feud.BaseBarony
}

var BRahachow罗加乔夫 feud.Barony = &罗加乔夫RahachowBarony{}

func init() {
	f := BRahachow罗加乔夫.(*罗加乔夫RahachowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rahachow",
		TitleName: "罗加乔夫",
		TitleCode: "b_rahachow",
	}
}
