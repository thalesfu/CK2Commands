package siuntio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛赫亚LohjaBarony struct {
	feud.BaseBarony
}

var BLohja洛赫亚 feud.Barony = &洛赫亚LohjaBarony{}

func init() {
    f := BLohja洛赫亚.(*洛赫亚LohjaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lohja",
		TitleName: "洛赫亚",
		TitleCode: "b_lohja",
	}
}
