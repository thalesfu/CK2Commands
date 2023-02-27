package sudovia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏瓦尔凯SuvalkaiBarony struct {
	feud.BaseBarony
}

var BSuvalkai苏瓦尔凯 feud.Barony = &苏瓦尔凯SuvalkaiBarony{}

func init() {
    f := BSuvalkai苏瓦尔凯.(*苏瓦尔凯SuvalkaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suvalkai",
		TitleName: "苏瓦尔凯",
		TitleCode: "b_suvalkai",
	}
}
