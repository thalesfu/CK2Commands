package kangly

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 康里KanglyBarony struct {
	feud.BaseBarony
}

var BKangly康里 feud.Barony = &康里KanglyBarony{}

func init() {
    f := BKangly康里.(*康里KanglyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kangly",
		TitleName: "康里",
		TitleCode: "b_kangly",
	}
}
