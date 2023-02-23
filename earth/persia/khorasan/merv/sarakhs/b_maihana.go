package sarakhs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈哈娜MaihanaBarony struct {
	feud.BaseBarony
}

var BMaihana迈哈娜 feud.Barony = &迈哈娜MaihanaBarony{}

func init() {
	f := BMaihana迈哈娜.(*迈哈娜MaihanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maihana",
		TitleName: "迈哈娜",
		TitleCode: "b_maihana",
	}
}
