package cieszyn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日维茨ZywiecBarony struct {
	feud.BaseBarony
}

var BZywiec日维茨 feud.Barony = &日维茨ZywiecBarony{}

func init() {
    f := BZywiec日维茨.(*日维茨ZywiecBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zywiec",
		TitleName: "日维茨",
		TitleCode: "b_zywiec",
	}
}
