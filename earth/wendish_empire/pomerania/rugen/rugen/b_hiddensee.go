package rugen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希登塞HiddenseeBarony struct {
	feud.BaseBarony
}

var BHiddensee希登塞 feud.Barony = &希登塞HiddenseeBarony{}

func init() {
    f := BHiddensee希登塞.(*希登塞HiddenseeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hiddensee",
		TitleName: "希登塞",
		TitleCode: "b_hiddensee",
	}
}
