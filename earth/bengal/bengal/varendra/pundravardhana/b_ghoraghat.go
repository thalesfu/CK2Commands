package pundravardhana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瞿罗伽吒GhoraghatBarony struct {
	feud.BaseBarony
}

var BGhoraghat瞿罗伽吒 feud.Barony = &瞿罗伽吒GhoraghatBarony{}

func init() {
	f := BGhoraghat瞿罗伽吒.(*瞿罗伽吒GhoraghatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghoraghat",
		TitleName: "瞿罗伽吒",
		TitleCode: "b_ghoraghat",
	}
}
