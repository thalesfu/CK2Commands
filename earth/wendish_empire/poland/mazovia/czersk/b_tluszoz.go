package czersk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特武什奇TluszozBarony struct {
	feud.BaseBarony
}

var BTluszoz特武什奇 feud.Barony = &特武什奇TluszozBarony{}

func init() {
    f := BTluszoz特武什奇.(*特武什奇TluszozBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tluszoz",
		TitleName: "特武什奇",
		TitleCode: "b_tluszoz",
	}
}
