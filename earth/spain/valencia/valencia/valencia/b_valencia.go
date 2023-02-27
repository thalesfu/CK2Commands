package valencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦伦西亚ValenciaBarony struct {
	feud.BaseBarony
}

var BValencia瓦伦西亚 feud.Barony = &瓦伦西亚ValenciaBarony{}

func init() {
    f := BValencia瓦伦西亚.(*瓦伦西亚ValenciaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valencia",
		TitleName: "瓦伦西亚",
		TitleCode: "b_valencia",
	}
}
