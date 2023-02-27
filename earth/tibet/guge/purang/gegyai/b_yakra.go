package gegyai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚热YakraBarony struct {
	feud.BaseBarony
}

var BYakra亚热 feud.Barony = &亚热YakraBarony{}

func init() {
    f := BYakra亚热.(*亚热YakraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yakra",
		TitleName: "亚热",
		TitleCode: "b_yakra",
	}
}
