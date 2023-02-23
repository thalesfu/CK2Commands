package korinthos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔戈斯ArgosBarony struct {
	feud.BaseBarony
}

var BArgos阿尔戈斯 feud.Barony = &阿尔戈斯ArgosBarony{}

func init() {
	f := BArgos阿尔戈斯.(*阿尔戈斯ArgosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "argos",
		TitleName: "阿尔戈斯",
		TitleCode: "b_argos",
	}
}
