package leon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦伦西亚德坎波斯ValenciadecamposBarony struct {
	feud.BaseBarony
}

var BValenciadecampos瓦伦西亚德坎波斯 feud.Barony = &瓦伦西亚德坎波斯ValenciadecamposBarony{}

func init() {
    f := BValenciadecampos瓦伦西亚德坎波斯.(*瓦伦西亚德坎波斯ValenciadecamposBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valenciadecampos",
		TitleName: "瓦伦西亚德坎波斯",
		TitleCode: "b_valenciadecampos",
	}
}
