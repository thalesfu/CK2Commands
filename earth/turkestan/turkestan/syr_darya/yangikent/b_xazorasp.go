package yangikent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎佐拉尔XazoraspBarony struct {
	feud.BaseBarony
}

var BXazorasp扎佐拉尔 feud.Barony = &扎佐拉尔XazoraspBarony{}

func init() {
    f := BXazorasp扎佐拉尔.(*扎佐拉尔XazoraspBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xazorasp",
		TitleName: "扎佐拉尔",
		TitleCode: "b_xazorasp",
	}
}
