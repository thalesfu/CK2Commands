package vijnot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科德吉GhotkiBarony struct {
	feud.BaseBarony
}

var BGhotki科德吉 feud.Barony = &科德吉GhotkiBarony{}

func init() {
    f := BGhotki科德吉.(*科德吉GhotkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghotki",
		TitleName: "科德吉",
		TitleCode: "b_ghotki",
	}
}
