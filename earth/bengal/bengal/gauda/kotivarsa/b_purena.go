package kotivarsa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富梨那PurenaBarony struct {
	feud.BaseBarony
}

var BPurena富梨那 feud.Barony = &富梨那PurenaBarony{}

func init() {
    f := BPurena富梨那.(*富梨那PurenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "purena",
		TitleName: "富梨那",
		TitleCode: "b_purena",
	}
}
