package soso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基里纳KirinaBarony struct {
	feud.BaseBarony
}

var BKirina基里纳 feud.Barony = &基里纳KirinaBarony{}

func init() {
    f := BKirina基里纳.(*基里纳KirinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirina",
		TitleName: "基里纳",
		TitleCode: "b_kirina",
	}
}
