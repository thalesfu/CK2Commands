package ikonion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 以哥念IkonionBarony struct {
	feud.BaseBarony
}

var BIkonion以哥念 feud.Barony = &以哥念IkonionBarony{}

func init() {
    f := BIkonion以哥念.(*以哥念IkonionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ikonion",
		TitleName: "以哥念",
		TitleCode: "b_ikonion",
	}
}
