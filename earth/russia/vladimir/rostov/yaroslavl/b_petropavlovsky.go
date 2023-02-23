package yaroslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彼得罗巴甫洛夫斯基PetropavlovskyBarony struct {
	feud.BaseBarony
}

var BPetropavlovsky彼得罗巴甫洛夫斯基 feud.Barony = &彼得罗巴甫洛夫斯基PetropavlovskyBarony{}

func init() {
	f := BPetropavlovsky彼得罗巴甫洛夫斯基.(*彼得罗巴甫洛夫斯基PetropavlovskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "petropavlovsky",
		TitleName: "彼得罗巴甫洛夫斯基",
		TitleCode: "b_petropavlovsky",
	}
}
