package tyrconnell

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法罕FahanBarony struct {
	feud.BaseBarony
}

var BFahan法罕 feud.Barony = &法罕FahanBarony{}

func init() {
	f := BFahan法罕.(*法罕FahanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fahan",
		TitleName: "法罕",
		TitleCode: "b_fahan",
	}
}
