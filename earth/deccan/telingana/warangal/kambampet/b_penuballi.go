package kambampet

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩努巴利PenuballiBarony struct {
	feud.BaseBarony
}

var BPenuballi佩努巴利 feud.Barony = &佩努巴利PenuballiBarony{}

func init() {
	f := BPenuballi佩努巴利.(*佩努巴利PenuballiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "penuballi",
		TitleName: "佩努巴利",
		TitleCode: "b_penuballi",
	}
}
