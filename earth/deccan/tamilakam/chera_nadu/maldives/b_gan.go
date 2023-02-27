package maldives

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 噶恩GanBarony struct {
	feud.BaseBarony
}

var BGan噶恩 feud.Barony = &噶恩GanBarony{}

func init() {
    f := BGan噶恩.(*噶恩GanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gan",
		TitleName: "噶恩",
		TitleCode: "b_gan",
	}
}
