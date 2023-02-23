package gegyai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 色喀执SekhadrigBarony struct {
	feud.BaseBarony
}

var BSekhadrig色喀执 feud.Barony = &色喀执SekhadrigBarony{}

func init() {
	f := BSekhadrig色喀执.(*色喀执SekhadrigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sekhadrig",
		TitleName: "色喀执",
		TitleCode: "b_sekhadrig",
	}
}
