package kanyakubja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 剑毕离KampilBarony struct {
	feud.BaseBarony
}

var BKampil剑毕离 feud.Barony = &剑毕离KampilBarony{}

func init() {
    f := BKampil剑毕离.(*剑毕离KampilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kampil",
		TitleName: "剑毕离",
		TitleCode: "b_kampil",
	}
}
