package idjil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁韦撒RouessaBarony struct {
	feud.BaseBarony
}

var BRouessa鲁韦撒 feud.Barony = &鲁韦撒RouessaBarony{}

func init() {
	f := BRouessa鲁韦撒.(*鲁韦撒RouessaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rouessa",
		TitleName: "鲁韦撒",
		TitleCode: "b_rouessa",
	}
}
