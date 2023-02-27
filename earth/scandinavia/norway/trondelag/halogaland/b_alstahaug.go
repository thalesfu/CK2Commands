package halogaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔斯塔海于格AlstahaugBarony struct {
	feud.BaseBarony
}

var BAlstahaug阿尔斯塔海于格 feud.Barony = &阿尔斯塔海于格AlstahaugBarony{}

func init() {
    f := BAlstahaug阿尔斯塔海于格.(*阿尔斯塔海于格AlstahaugBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alstahaug",
		TitleName: "阿尔斯塔海于格",
		TitleCode: "b_alstahaug",
	}
}
