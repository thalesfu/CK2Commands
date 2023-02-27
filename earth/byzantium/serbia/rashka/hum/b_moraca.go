package hum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫拉查MoracaBarony struct {
	feud.BaseBarony
}

var BMoraca莫拉查 feud.Barony = &莫拉查MoracaBarony{}

func init() {
    f := BMoraca莫拉查.(*莫拉查MoracaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moraca",
		TitleName: "莫拉查",
		TitleCode: "b_moraca",
	}
}
