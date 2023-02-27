package sarasvata_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘波迦KamboikaBarony struct {
	feud.BaseBarony
}

var BKamboika甘波迦 feud.Barony = &甘波迦KamboikaBarony{}

func init() {
    f := BKamboika甘波迦.(*甘波迦KamboikaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamboika",
		TitleName: "甘波迦",
		TitleCode: "b_kamboika",
	}
}
