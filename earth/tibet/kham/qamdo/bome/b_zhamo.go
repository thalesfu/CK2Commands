package bome

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎木ZhamoBarony struct {
	feud.BaseBarony
}

var BZhamo扎木 feud.Barony = &扎木ZhamoBarony{}

func init() {
    f := BZhamo扎木.(*扎木ZhamoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhamo",
		TitleName: "扎木",
		TitleCode: "b_zhamo",
	}
}
