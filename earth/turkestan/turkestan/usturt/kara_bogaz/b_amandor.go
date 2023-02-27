package kara_bogaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿曼多尔AmandorBarony struct {
	feud.BaseBarony
}

var BAmandor阿曼多尔 feud.Barony = &阿曼多尔AmandorBarony{}

func init() {
    f := BAmandor阿曼多尔.(*阿曼多尔AmandorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amandor",
		TitleName: "阿曼多尔",
		TitleCode: "b_amandor",
	}
}
