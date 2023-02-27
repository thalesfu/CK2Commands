package munda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩摩卢顿MelmarudunBarony struct {
	feud.BaseBarony
}

var BMelmarudun摩摩卢顿 feud.Barony = &摩摩卢顿MelmarudunBarony{}

func init() {
    f := BMelmarudun摩摩卢顿.(*摩摩卢顿MelmarudunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "melmarudun",
		TitleName: "摩摩卢顿",
		TitleCode: "b_melmarudun",
	}
}
