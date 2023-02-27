package krain

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施泰因StainBarony struct {
	feud.BaseBarony
}

var BStain施泰因 feud.Barony = &施泰因StainBarony{}

func init() {
    f := BStain施泰因.(*施泰因StainBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stain",
		TitleName: "施泰因",
		TitleCode: "b_stain",
	}
}
