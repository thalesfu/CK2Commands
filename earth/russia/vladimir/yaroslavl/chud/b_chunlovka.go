package chud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 春洛夫卡ChunlovkaBarony struct {
	feud.BaseBarony
}

var BChunlovka春洛夫卡 feud.Barony = &春洛夫卡ChunlovkaBarony{}

func init() {
    f := BChunlovka春洛夫卡.(*春洛夫卡ChunlovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chunlovka",
		TitleName: "春洛夫卡",
		TitleCode: "b_chunlovka",
	}
}
