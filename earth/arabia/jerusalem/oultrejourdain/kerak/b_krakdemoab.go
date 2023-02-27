package kerak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩押的吉珥KrakdemoabBarony struct {
	feud.BaseBarony
}

var BKrakdemoab摩押的吉珥 feud.Barony = &摩押的吉珥KrakdemoabBarony{}

func init() {
    f := BKrakdemoab摩押的吉珥.(*摩押的吉珥KrakdemoabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krakdemoab",
		TitleName: "摩押的吉珥",
		TitleCode: "b_krakdemoab",
	}
}
