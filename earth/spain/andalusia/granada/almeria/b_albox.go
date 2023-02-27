package almeria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔沃克斯AlboxBarony struct {
	feud.BaseBarony
}

var BAlbox阿尔沃克斯 feud.Barony = &阿尔沃克斯AlboxBarony{}

func init() {
    f := BAlbox阿尔沃克斯.(*阿尔沃克斯AlboxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "albox",
		TitleName: "阿尔沃克斯",
		TitleCode: "b_albox",
	}
}
