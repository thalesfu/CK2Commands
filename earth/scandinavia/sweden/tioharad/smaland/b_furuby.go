package smaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富鲁比FurubyBarony struct {
	feud.BaseBarony
}

var BFuruby富鲁比 feud.Barony = &富鲁比FurubyBarony{}

func init() {
	f := BFuruby富鲁比.(*富鲁比FurubyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "furuby",
		TitleName: "富鲁比",
		TitleCode: "b_furuby",
	}
}
