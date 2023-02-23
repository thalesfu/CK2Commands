package onega

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彼得罗斯科伊PetrozavodskBarony struct {
	feud.BaseBarony
}

var BPetrozavodsk彼得罗斯科伊 feud.Barony = &彼得罗斯科伊PetrozavodskBarony{}

func init() {
	f := BPetrozavodsk彼得罗斯科伊.(*彼得罗斯科伊PetrozavodskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "petrozavodsk",
		TitleName: "彼得罗斯科伊",
		TitleCode: "b_petrozavodsk",
	}
}
