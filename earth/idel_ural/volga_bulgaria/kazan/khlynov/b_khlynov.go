package khlynov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫雷诺夫KhlynovBarony struct {
	feud.BaseBarony
}

var BKhlynov赫雷诺夫 feud.Barony = &赫雷诺夫KhlynovBarony{}

func init() {
    f := BKhlynov赫雷诺夫.(*赫雷诺夫KhlynovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khlynov",
		TitleName: "赫雷诺夫",
		TitleCode: "b_khlynov",
	}
}
