package kongu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂卢补罗TiruppurBarony struct {
	feud.BaseBarony
}

var BTiruppur蒂卢补罗 feud.Barony = &蒂卢补罗TiruppurBarony{}

func init() {
    f := BTiruppur蒂卢补罗.(*蒂卢补罗TiruppurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiruppur",
		TitleName: "蒂卢补罗",
		TitleCode: "b_tiruppur",
	}
}
