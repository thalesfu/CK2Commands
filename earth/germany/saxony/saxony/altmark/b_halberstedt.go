package altmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈尔伯施塔特HalberstedtBarony struct {
	feud.BaseBarony
}

var BHalberstedt哈尔伯施塔特 feud.Barony = &哈尔伯施塔特HalberstedtBarony{}

func init() {
    f := BHalberstedt哈尔伯施塔特.(*哈尔伯施塔特HalberstedtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "halberstedt",
		TitleName: "哈尔伯施塔特",
		TitleCode: "b_halberstedt",
	}
}
