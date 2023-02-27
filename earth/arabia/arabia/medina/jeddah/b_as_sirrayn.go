package jeddah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡赖因As_sirraynBarony struct {
	feud.BaseBarony
}

var BAs_sirrayn锡赖因 feud.Barony = &锡赖因As_sirraynBarony{}

func init() {
    f := BAs_sirrayn锡赖因.(*锡赖因As_sirraynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "as_sirrayn",
		TitleName: "锡赖因",
		TitleCode: "b_as_sirrayn",
	}
}
