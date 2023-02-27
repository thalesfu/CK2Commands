package bukhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰什肯特VabkentBarony struct {
	feud.BaseBarony
}

var BVabkent恰什肯特 feud.Barony = &恰什肯特VabkentBarony{}

func init() {
    f := BVabkent恰什肯特.(*恰什肯特VabkentBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vabkent",
		TitleName: "恰什肯特",
		TitleCode: "b_vabkent",
	}
}
