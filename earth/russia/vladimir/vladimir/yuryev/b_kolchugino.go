package yuryev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科利丘吉诺KolchuginoBarony struct {
	feud.BaseBarony
}

var BKolchugino科利丘吉诺 feud.Barony = &科利丘吉诺KolchuginoBarony{}

func init() {
    f := BKolchugino科利丘吉诺.(*科利丘吉诺KolchuginoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolchugino",
		TitleName: "科利丘吉诺",
		TitleCode: "b_kolchugino",
	}
}
