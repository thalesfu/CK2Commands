package maan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿夫拉AfraBarony struct {
	feud.BaseBarony
}

var BAfra阿夫拉 feud.Barony = &阿夫拉AfraBarony{}

func init() {
    f := BAfra阿夫拉.(*阿夫拉AfraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "afra",
		TitleName: "阿夫拉",
		TitleCode: "b_afra",
	}
}
