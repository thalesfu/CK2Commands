package santiago

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蓬特韦德拉PontevedraBarony struct {
	feud.BaseBarony
}

var BPontevedra蓬特韦德拉 feud.Barony = &蓬特韦德拉PontevedraBarony{}

func init() {
    f := BPontevedra蓬特韦德拉.(*蓬特韦德拉PontevedraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pontevedra",
		TitleName: "蓬特韦德拉",
		TitleCode: "b_pontevedra",
	}
}
