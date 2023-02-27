package canda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旃陀CandaBarony struct {
	feud.BaseBarony
}

var BCanda旃陀 feud.Barony = &旃陀CandaBarony{}

func init() {
    f := BCanda旃陀.(*旃陀CandaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "canda",
		TitleName: "旃陀",
		TitleCode: "b_canda",
	}
}
