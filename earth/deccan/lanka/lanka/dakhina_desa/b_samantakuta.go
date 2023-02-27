package dakhina_desa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 三曼多矩吒SamantakutaBarony struct {
	feud.BaseBarony
}

var BSamantakuta三曼多矩吒 feud.Barony = &三曼多矩吒SamantakutaBarony{}

func init() {
    f := BSamantakuta三曼多矩吒.(*三曼多矩吒SamantakutaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samantakuta",
		TitleName: "三曼多矩吒",
		TitleCode: "b_samantakuta",
	}
}
