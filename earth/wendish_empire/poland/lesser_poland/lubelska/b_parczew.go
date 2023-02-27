package lubelska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕尔切夫ParczewBarony struct {
	feud.BaseBarony
}

var BParczew帕尔切夫 feud.Barony = &帕尔切夫ParczewBarony{}

func init() {
    f := BParczew帕尔切夫.(*帕尔切夫ParczewBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "parczew",
		TitleName: "帕尔切夫",
		TitleCode: "b_parczew",
	}
}
