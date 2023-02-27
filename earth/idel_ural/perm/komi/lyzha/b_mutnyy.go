package lyzha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆特内MutnyyBarony struct {
	feud.BaseBarony
}

var BMutnyy穆特内 feud.Barony = &穆特内MutnyyBarony{}

func init() {
    f := BMutnyy穆特内.(*穆特内MutnyyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mutnyy",
		TitleName: "穆特内",
		TitleCode: "b_mutnyy",
	}
}
