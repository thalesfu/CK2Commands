package yolva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科因KoinBarony struct {
	feud.BaseBarony
}

var BKoin科因 feud.Barony = &科因KoinBarony{}

func init() {
    f := BKoin科因.(*科因KoinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koin",
		TitleName: "科因",
		TitleCode: "b_koin",
	}
}
