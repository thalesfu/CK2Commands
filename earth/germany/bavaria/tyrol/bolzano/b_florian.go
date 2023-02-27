package bolzano

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗洛里安FlorianBarony struct {
	feud.BaseBarony
}

var BFlorian弗洛里安 feud.Barony = &弗洛里安FlorianBarony{}

func init() {
    f := BFlorian弗洛里安.(*弗洛里安FlorianBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "florian",
		TitleName: "弗洛里安",
		TitleCode: "b_florian",
	}
}
