package starodub

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿夫迪夫卡AvdiikaBarony struct {
	feud.BaseBarony
}

var BAvdiika阿夫迪夫卡 feud.Barony = &阿夫迪夫卡AvdiikaBarony{}

func init() {
    f := BAvdiika阿夫迪夫卡.(*阿夫迪夫卡AvdiikaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avdiika",
		TitleName: "阿夫迪夫卡",
		TitleCode: "b_avdiika",
	}
}
