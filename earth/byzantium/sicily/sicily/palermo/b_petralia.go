package palermo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩特拉利亚PetraliaBarony struct {
	feud.BaseBarony
}

var BPetralia佩特拉利亚 feud.Barony = &佩特拉利亚PetraliaBarony{}

func init() {
    f := BPetralia佩特拉利亚.(*佩特拉利亚PetraliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "petralia",
		TitleName: "佩特拉利亚",
		TitleCode: "b_petralia",
	}
}
