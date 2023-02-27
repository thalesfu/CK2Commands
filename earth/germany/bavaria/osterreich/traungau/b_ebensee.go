package traungau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃本塞EbenseeBarony struct {
	feud.BaseBarony
}

var BEbensee埃本塞 feud.Barony = &埃本塞EbenseeBarony{}

func init() {
    f := BEbensee埃本塞.(*埃本塞EbenseeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ebensee",
		TitleName: "埃本塞",
		TitleCode: "b_ebensee",
	}
}
