package aoudaghost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯达马KedamaBarony struct {
	feud.BaseBarony
}

var BKedama凯达马 feud.Barony = &凯达马KedamaBarony{}

func init() {
	f := BKedama凯达马.(*凯达马KedamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kedama",
		TitleName: "凯达马",
		TitleCode: "b_kedama",
	}
}
