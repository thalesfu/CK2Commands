package jacwiez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯基达利SkidzielBarony struct {
	feud.BaseBarony
}

var BSkidziel斯基达利 feud.Barony = &斯基达利SkidzielBarony{}

func init() {
    f := BSkidziel斯基达利.(*斯基达利SkidzielBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skidziel",
		TitleName: "斯基达利",
		TitleCode: "b_skidziel",
	}
}
