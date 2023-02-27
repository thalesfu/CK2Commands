package talava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿卢克斯内AluksneBarony struct {
	feud.BaseBarony
}

var BAluksne阿卢克斯内 feud.Barony = &阿卢克斯内AluksneBarony{}

func init() {
    f := BAluksne阿卢克斯内.(*阿卢克斯内AluksneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aluksne",
		TitleName: "阿卢克斯内",
		TitleCode: "b_aluksne",
	}
}
