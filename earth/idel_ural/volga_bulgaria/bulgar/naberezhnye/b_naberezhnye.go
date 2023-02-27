package naberezhnye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳别列日内耶NaberezhnyeBarony struct {
	feud.BaseBarony
}

var BNaberezhnye纳别列日内耶 feud.Barony = &纳别列日内耶NaberezhnyeBarony{}

func init() {
    f := BNaberezhnye纳别列日内耶.(*纳别列日内耶NaberezhnyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naberezhnye",
		TitleName: "纳别列日内耶",
		TitleCode: "b_naberezhnye",
	}
}
