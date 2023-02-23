package asni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 牟沙那揭罗MusanagarBarony struct {
	feud.BaseBarony
}

var BMusanagar牟沙那揭罗 feud.Barony = &牟沙那揭罗MusanagarBarony{}

func init() {
	f := BMusanagar牟沙那揭罗.(*牟沙那揭罗MusanagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "musanagar",
		TitleName: "牟沙那揭罗",
		TitleCode: "b_musanagar",
	}
}
