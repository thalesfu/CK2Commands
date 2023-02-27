package gao

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加奥_撒内GaosaneyBarony struct {
	feud.BaseBarony
}

var BGaosaney加奥_撒内 feud.Barony = &加奥_撒内GaosaneyBarony{}

func init() {
    f := BGaosaney加奥_撒内.(*加奥_撒内GaosaneyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gaosaney",
		TitleName: "加奥_撒内",
		TitleCode: "b_gaosaney",
	}
}
