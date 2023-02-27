package leh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉隆KeylongBarony struct {
	feud.BaseBarony
}

var BKeylong吉隆 feud.Barony = &吉隆KeylongBarony{}

func init() {
    f := BKeylong吉隆.(*吉隆KeylongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "keylong",
		TitleName: "吉隆",
		TitleCode: "b_keylong",
	}
}
