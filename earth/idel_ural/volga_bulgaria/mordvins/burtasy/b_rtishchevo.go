package burtasy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒季谢沃RtishchevoBarony struct {
	feud.BaseBarony
}

var BRtishchevo勒季谢沃 feud.Barony = &勒季谢沃RtishchevoBarony{}

func init() {
    f := BRtishchevo勒季谢沃.(*勒季谢沃RtishchevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rtishchevo",
		TitleName: "勒季谢沃",
		TitleCode: "b_rtishchevo",
	}
}
