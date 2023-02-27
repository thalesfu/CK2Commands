package borovichi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基里奇KirichiBarony struct {
	feud.BaseBarony
}

var BKirichi基里奇 feud.Barony = &基里奇KirichiBarony{}

func init() {
    f := BKirichi基里奇.(*基里奇KirichiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirichi",
		TitleName: "基里奇",
		TitleCode: "b_kirichi",
	}
}
