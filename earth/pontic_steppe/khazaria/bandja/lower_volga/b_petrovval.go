package lower_volga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彼得罗夫瓦尔PetrovvalBarony struct {
	feud.BaseBarony
}

var BPetrovval彼得罗夫瓦尔 feud.Barony = &彼得罗夫瓦尔PetrovvalBarony{}

func init() {
    f := BPetrovval彼得罗夫瓦尔.(*彼得罗夫瓦尔PetrovvalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "petrovval",
		TitleName: "彼得罗夫瓦尔",
		TitleCode: "b_petrovval",
	}
}
