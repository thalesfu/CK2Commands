package jarva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伦波伊斯AlempoisBarony struct {
	feud.BaseBarony
}

var BAlempois阿伦波伊斯 feud.Barony = &阿伦波伊斯AlempoisBarony{}

func init() {
    f := BAlempois阿伦波伊斯.(*阿伦波伊斯AlempoisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alempois",
		TitleName: "阿伦波伊斯",
		TitleCode: "b_alempois",
	}
}
