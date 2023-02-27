package kafirkot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德拉加济汗Dera_ghazi_khanBarony struct {
	feud.BaseBarony
}

var BDera_ghazi_khan德拉加济汗 feud.Barony = &德拉加济汗Dera_ghazi_khanBarony{}

func init() {
    f := BDera_ghazi_khan德拉加济汗.(*德拉加济汗Dera_ghazi_khanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dera_ghazi_khan",
		TitleName: "德拉加济汗",
		TitleCode: "b_dera_ghazi_khan",
	}
}
