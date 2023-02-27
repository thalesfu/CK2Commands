package buchan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因弗鲁里InverurieBarony struct {
	feud.BaseBarony
}

var BInverurie因弗鲁里 feud.Barony = &因弗鲁里InverurieBarony{}

func init() {
    f := BInverurie因弗鲁里.(*因弗鲁里InverurieBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "inverurie",
		TitleName: "因弗鲁里",
		TitleCode: "b_inverurie",
	}
}
