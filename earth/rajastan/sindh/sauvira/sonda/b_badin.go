package sonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆陈BadinBarony struct {
	feud.BaseBarony
}

var BBadin婆陈 feud.Barony = &婆陈BadinBarony{}

func init() {
	f := BBadin婆陈.(*婆陈BadinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "badin",
		TitleName: "婆陈",
		TitleCode: "b_badin",
	}
}
