package ladistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洪季KhonjBarony struct {
	feud.BaseBarony
}

var BKhonj洪季 feud.Barony = &洪季KhonjBarony{}

func init() {
	f := BKhonj洪季.(*洪季KhonjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khonj",
		TitleName: "洪季",
		TitleCode: "b_khonj",
	}
}
