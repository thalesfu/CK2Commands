package varanasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鹿野苑SarnathBarony struct {
	feud.BaseBarony
}

var BSarnath鹿野苑 feud.Barony = &鹿野苑SarnathBarony{}

func init() {
	f := BSarnath鹿野苑.(*鹿野苑SarnathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarnath",
		TitleName: "鹿野苑",
		TitleCode: "b_sarnath",
	}
}
