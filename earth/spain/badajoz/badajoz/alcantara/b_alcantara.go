package alcantara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔坎塔拉AlcantaraBarony struct {
	feud.BaseBarony
}

var BAlcantara阿尔坎塔拉 feud.Barony = &阿尔坎塔拉AlcantaraBarony{}

func init() {
    f := BAlcantara阿尔坎塔拉.(*阿尔坎塔拉AlcantaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alcantara",
		TitleName: "阿尔坎塔拉",
		TitleCode: "b_alcantara",
	}
}
