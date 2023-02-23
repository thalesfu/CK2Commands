package amman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨勒特SaltBarony struct {
	feud.BaseBarony
}

var BSalt萨勒特 feud.Barony = &萨勒特SaltBarony{}

func init() {
	f := BSalt萨勒特.(*萨勒特SaltBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salt",
		TitleName: "萨勒特",
		TitleCode: "b_salt",
	}
}
