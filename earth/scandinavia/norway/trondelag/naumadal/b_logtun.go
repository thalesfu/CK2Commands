package naumadal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗格顿LogtunBarony struct {
	feud.BaseBarony
}

var BLogtun罗格顿 feud.Barony = &罗格顿LogtunBarony{}

func init() {
    f := BLogtun罗格顿.(*罗格顿LogtunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "logtun",
		TitleName: "罗格顿",
		TitleCode: "b_logtun",
	}
}
