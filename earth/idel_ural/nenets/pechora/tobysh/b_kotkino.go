package tobysh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科特基诺KotkinoBarony struct {
	feud.BaseBarony
}

var BKotkino科特基诺 feud.Barony = &科特基诺KotkinoBarony{}

func init() {
    f := BKotkino科特基诺.(*科特基诺KotkinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kotkino",
		TitleName: "科特基诺",
		TitleCode: "b_kotkino",
	}
}
