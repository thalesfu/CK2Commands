package mahra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基什恩QishnBarony struct {
	feud.BaseBarony
}

var BQishn基什恩 feud.Barony = &基什恩QishnBarony{}

func init() {
    f := BQishn基什恩.(*基什恩QishnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qishn",
		TitleName: "基什恩",
		TitleCode: "b_qishn",
	}
}
