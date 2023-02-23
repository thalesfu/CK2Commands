package claudiopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 派特布里格PetobrigaBarony struct {
	feud.BaseBarony
}

var BPetobriga派特布里格 feud.Barony = &派特布里格PetobrigaBarony{}

func init() {
	f := BPetobriga派特布里格.(*派特布里格PetobrigaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "petobriga",
		TitleName: "派特布里格",
		TitleCode: "b_petobriga",
	}
}
