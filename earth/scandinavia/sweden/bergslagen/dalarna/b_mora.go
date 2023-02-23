package dalarna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆拉MoraBarony struct {
	feud.BaseBarony
}

var BMora穆拉 feud.Barony = &穆拉MoraBarony{}

func init() {
	f := BMora穆拉.(*穆拉MoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mora",
		TitleName: "穆拉",
		TitleCode: "b_mora",
	}
}
