package beni_yanni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希泽ThizaBarony struct {
	feud.BaseBarony
}

var BThiza希泽 feud.Barony = &希泽ThizaBarony{}

func init() {
    f := BThiza希泽.(*希泽ThizaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thiza",
		TitleName: "希泽",
		TitleCode: "b_thiza",
	}
}
