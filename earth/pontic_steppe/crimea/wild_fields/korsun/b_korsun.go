package korsun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔孙KorsunBarony struct {
	feud.BaseBarony
}

var BKorsun科尔孙 feud.Barony = &科尔孙KorsunBarony{}

func init() {
    f := BKorsun科尔孙.(*科尔孙KorsunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korsun",
		TitleName: "科尔孙",
		TitleCode: "b_korsun",
	}
}
