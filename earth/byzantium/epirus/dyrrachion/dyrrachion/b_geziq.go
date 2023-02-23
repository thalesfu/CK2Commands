package dyrrachion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格齐奇GeziqBarony struct {
	feud.BaseBarony
}

var BGeziq格齐奇 feud.Barony = &格齐奇GeziqBarony{}

func init() {
	f := BGeziq格齐奇.(*格齐奇GeziqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "geziq",
		TitleName: "格齐奇",
		TitleCode: "b_geziq",
	}
}
