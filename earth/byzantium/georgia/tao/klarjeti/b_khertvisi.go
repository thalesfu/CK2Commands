package klarjeti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫尔特维西KhertvisiBarony struct {
	feud.BaseBarony
}

var BKhertvisi赫尔特维西 feud.Barony = &赫尔特维西KhertvisiBarony{}

func init() {
    f := BKhertvisi赫尔特维西.(*赫尔特维西KhertvisiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khertvisi",
		TitleName: "赫尔特维西",
		TitleCode: "b_khertvisi",
	}
}
