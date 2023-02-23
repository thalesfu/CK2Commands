package nagchu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔茫KhormangBarony struct {
	feud.BaseBarony
}

var BKhormang库尔茫 feud.Barony = &库尔茫KhormangBarony{}

func init() {
	f := BKhormang库尔茫.(*库尔茫KhormangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khormang",
		TitleName: "库尔茫",
		TitleCode: "b_khormang",
	}
}
