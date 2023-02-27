package eichstadt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 魏森堡WeissenburgBarony struct {
	feud.BaseBarony
}

var BWeissenburg魏森堡 feud.Barony = &魏森堡WeissenburgBarony{}

func init() {
    f := BWeissenburg魏森堡.(*魏森堡WeissenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "weissenburg",
		TitleName: "魏森堡",
		TitleCode: "b_weissenburg",
	}
}
