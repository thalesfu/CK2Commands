package ephesos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕特仑PetronBarony struct {
	feud.BaseBarony
}

var BPetron帕特仑 feud.Barony = &帕特仑PetronBarony{}

func init() {
	f := BPetron帕特仑.(*帕特仑PetronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "petron",
		TitleName: "帕特仑",
		TitleCode: "b_petron",
	}
}
