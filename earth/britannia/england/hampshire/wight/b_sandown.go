package wight

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑当SandownBarony struct {
	feud.BaseBarony
}

var BSandown桑当 feud.Barony = &桑当SandownBarony{}

func init() {
    f := BSandown桑当.(*桑当SandownBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sandown",
		TitleName: "桑当",
		TitleCode: "b_sandown",
	}
}
