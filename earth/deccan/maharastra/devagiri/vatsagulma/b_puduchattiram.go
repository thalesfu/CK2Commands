package vatsagulma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富菟遮底蓝PuduchattiramBarony struct {
	feud.BaseBarony
}

var BPuduchattiram富菟遮底蓝 feud.Barony = &富菟遮底蓝PuduchattiramBarony{}

func init() {
	f := BPuduchattiram富菟遮底蓝.(*富菟遮底蓝PuduchattiramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "puduchattiram",
		TitleName: "富菟遮底蓝",
		TitleCode: "b_puduchattiram",
	}
}
