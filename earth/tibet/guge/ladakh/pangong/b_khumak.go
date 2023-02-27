package pangong

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔纳克KhumakBarony struct {
	feud.BaseBarony
}

var BKhumak库尔纳克 feud.Barony = &库尔纳克KhumakBarony{}

func init() {
    f := BKhumak库尔纳克.(*库尔纳克KhumakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khumak",
		TitleName: "库尔纳克",
		TitleCode: "b_khumak",
	}
}
