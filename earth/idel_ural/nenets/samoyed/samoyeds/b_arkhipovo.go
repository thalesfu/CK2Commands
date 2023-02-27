package samoyeds

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔希波沃ArkhipovoBarony struct {
	feud.BaseBarony
}

var BArkhipovo阿尔希波沃 feud.Barony = &阿尔希波沃ArkhipovoBarony{}

func init() {
    f := BArkhipovo阿尔希波沃.(*阿尔希波沃ArkhipovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arkhipovo",
		TitleName: "阿尔希波沃",
		TitleCode: "b_arkhipovo",
	}
}
