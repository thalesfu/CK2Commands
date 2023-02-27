package sonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦卢恩奢KaroonjarBarony struct {
	feud.BaseBarony
}

var BKaroonjar迦卢恩奢 feud.Barony = &迦卢恩奢KaroonjarBarony{}

func init() {
    f := BKaroonjar迦卢恩奢.(*迦卢恩奢KaroonjarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karoonjar",
		TitleName: "迦卢恩奢",
		TitleCode: "b_karoonjar",
	}
}
