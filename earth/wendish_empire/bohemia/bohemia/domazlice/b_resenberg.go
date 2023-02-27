package domazlice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷森贝格ResenbergBarony struct {
	feud.BaseBarony
}

var BResenberg雷森贝格 feud.Barony = &雷森贝格ResenbergBarony{}

func init() {
    f := BResenberg雷森贝格.(*雷森贝格ResenbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "resenberg",
		TitleName: "雷森贝格",
		TitleCode: "b_resenberg",
	}
}
