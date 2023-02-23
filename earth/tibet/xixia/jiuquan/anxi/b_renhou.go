package anxi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 仁厚RenhouBarony struct {
	feud.BaseBarony
}

var BRenhou仁厚 feud.Barony = &仁厚RenhouBarony{}

func init() {
	f := BRenhou仁厚.(*仁厚RenhouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "renhou",
		TitleName: "仁厚",
		TitleCode: "b_renhou",
	}
}
