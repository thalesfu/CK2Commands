package cherson

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫尔松KhersonBarony struct {
	feud.BaseBarony
}

var BKherson赫尔松 feud.Barony = &赫尔松KhersonBarony{}

func init() {
    f := BKherson赫尔松.(*赫尔松KhersonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kherson",
		TitleName: "赫尔松",
		TitleCode: "b_kherson",
	}
}
