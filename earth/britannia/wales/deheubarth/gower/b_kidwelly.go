package gower

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基德韦利KidwellyBarony struct {
	feud.BaseBarony
}

var BKidwelly基德韦利 feud.Barony = &基德韦利KidwellyBarony{}

func init() {
	f := BKidwelly基德韦利.(*基德韦利KidwellyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kidwelly",
		TitleName: "基德韦利",
		TitleCode: "b_kidwelly",
	}
}
