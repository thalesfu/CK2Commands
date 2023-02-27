package philippopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彼得里奇PetrichBarony struct {
	feud.BaseBarony
}

var BPetrich彼得里奇 feud.Barony = &彼得里奇PetrichBarony{}

func init() {
    f := BPetrich彼得里奇.(*彼得里奇PetrichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "petrich",
		TitleName: "彼得里奇",
		TitleCode: "b_petrich",
	}
}
