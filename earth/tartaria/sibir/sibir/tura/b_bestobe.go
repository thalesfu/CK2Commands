package tura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别斯托别BestobeBarony struct {
	feud.BaseBarony
}

var BBestobe别斯托别 feud.Barony = &别斯托别BestobeBarony{}

func init() {
	f := BBestobe别斯托别.(*别斯托别BestobeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bestobe",
		TitleName: "别斯托别",
		TitleCode: "b_bestobe",
	}
}
