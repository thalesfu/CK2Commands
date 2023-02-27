package aulon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索波特Sopot_aulonBarony struct {
	feud.BaseBarony
}

var BSopot_aulon索波特 feud.Barony = &索波特Sopot_aulonBarony{}

func init() {
    f := BSopot_aulon索波特.(*索波特Sopot_aulonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sopot_aulon",
		TitleName: "索波特",
		TitleCode: "b_sopot_aulon",
	}
}
