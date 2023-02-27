package yolva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维姆VymBarony struct {
	feud.BaseBarony
}

var BVym维姆 feud.Barony = &维姆VymBarony{}

func init() {
    f := BVym维姆.(*维姆VymBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vym",
		TitleName: "维姆",
		TitleCode: "b_vym",
	}
}
