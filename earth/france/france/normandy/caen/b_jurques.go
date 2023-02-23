package caen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑞尔克JurquesBarony struct {
	feud.BaseBarony
}

var BJurques瑞尔克 feud.Barony = &瑞尔克JurquesBarony{}

func init() {
	f := BJurques瑞尔克.(*瑞尔克JurquesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jurques",
		TitleName: "瑞尔克",
		TitleCode: "b_jurques",
	}
}
