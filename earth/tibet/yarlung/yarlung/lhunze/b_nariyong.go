package lhunze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拿日雍NariyongBarony struct {
	feud.BaseBarony
}

var BNariyong拿日雍 feud.Barony = &拿日雍NariyongBarony{}

func init() {
    f := BNariyong拿日雍.(*拿日雍NariyongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nariyong",
		TitleName: "拿日雍",
		TitleCode: "b_nariyong",
	}
}
