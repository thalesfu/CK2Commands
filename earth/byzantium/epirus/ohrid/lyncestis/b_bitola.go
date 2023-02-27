package lyncestis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比托拉BitolaBarony struct {
	feud.BaseBarony
}

var BBitola比托拉 feud.Barony = &比托拉BitolaBarony{}

func init() {
    f := BBitola比托拉.(*比托拉BitolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bitola",
		TitleName: "比托拉",
		TitleCode: "b_bitola",
	}
}
