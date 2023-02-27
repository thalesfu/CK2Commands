package vasyugan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳济诺NazinoBarony struct {
	feud.BaseBarony
}

var BNazino纳济诺 feud.Barony = &纳济诺NazinoBarony{}

func init() {
    f := BNazino纳济诺.(*纳济诺NazinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nazino",
		TitleName: "纳济诺",
		TitleCode: "b_nazino",
	}
}
