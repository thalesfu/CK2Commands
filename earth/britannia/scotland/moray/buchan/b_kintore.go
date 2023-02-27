package buchan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 金托尔KintoreBarony struct {
	feud.BaseBarony
}

var BKintore金托尔 feud.Barony = &金托尔KintoreBarony{}

func init() {
    f := BKintore金托尔.(*金托尔KintoreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kintore",
		TitleName: "金托尔",
		TitleCode: "b_kintore",
	}
}
