package poznanskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 缅济热奇MiedzyrzeczBarony struct {
	feud.BaseBarony
}

var BMiedzyrzecz缅济热奇 feud.Barony = &缅济热奇MiedzyrzeczBarony{}

func init() {
    f := BMiedzyrzecz缅济热奇.(*缅济热奇MiedzyrzeczBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "miedzyrzecz",
		TitleName: "缅济热奇",
		TitleCode: "b_miedzyrzecz",
	}
}
