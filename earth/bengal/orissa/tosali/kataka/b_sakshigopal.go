package kataka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索崎瞿波罗SakshigopalBarony struct {
	feud.BaseBarony
}

var BSakshigopal索崎瞿波罗 feud.Barony = &索崎瞿波罗SakshigopalBarony{}

func init() {
    f := BSakshigopal索崎瞿波罗.(*索崎瞿波罗SakshigopalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sakshigopal",
		TitleName: "索崎瞿波罗",
		TitleCode: "b_sakshigopal",
	}
}
