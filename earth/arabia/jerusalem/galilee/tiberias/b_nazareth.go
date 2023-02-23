package tiberias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拿撒勒NazarethBarony struct {
	feud.BaseBarony
}

var BNazareth拿撒勒 feud.Barony = &拿撒勒NazarethBarony{}

func init() {
	f := BNazareth拿撒勒.(*拿撒勒NazarethBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nazareth",
		TitleName: "拿撒勒",
		TitleCode: "b_nazareth",
	}
}
