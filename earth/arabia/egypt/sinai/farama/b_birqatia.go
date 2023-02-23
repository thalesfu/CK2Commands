package farama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖提耶井BirqatiaBarony struct {
	feud.BaseBarony
}

var BBirqatia盖提耶井 feud.Barony = &盖提耶井BirqatiaBarony{}

func init() {
	f := BBirqatia盖提耶井.(*盖提耶井BirqatiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "birqatia",
		TitleName: "盖提耶井",
		TitleCode: "b_birqatia",
	}
}
