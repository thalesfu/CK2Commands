package tavasts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海门林纳HameenlinnaBarony struct {
	feud.BaseBarony
}

var BHameenlinna海门林纳 feud.Barony = &海门林纳HameenlinnaBarony{}

func init() {
	f := BHameenlinna海门林纳.(*海门林纳HameenlinnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hameenlinna",
		TitleName: "海门林纳",
		TitleCode: "b_hameenlinna",
	}
}
