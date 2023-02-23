package udabhanda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥希德OhindBarony struct {
	feud.BaseBarony
}

var BOhind奥希德 feud.Barony = &奥希德OhindBarony{}

func init() {
	f := BOhind奥希德.(*奥希德OhindBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ohind",
		TitleName: "奥希德",
		TitleCode: "b_ohind",
	}
}
