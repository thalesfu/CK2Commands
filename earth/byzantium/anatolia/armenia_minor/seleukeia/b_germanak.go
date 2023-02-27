package seleukeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日耳曼纳克GermanakBarony struct {
	feud.BaseBarony
}

var BGermanak日耳曼纳克 feud.Barony = &日耳曼纳克GermanakBarony{}

func init() {
    f := BGermanak日耳曼纳克.(*日耳曼纳克GermanakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "germanak",
		TitleName: "日耳曼纳克",
		TitleCode: "b_germanak",
	}
}
