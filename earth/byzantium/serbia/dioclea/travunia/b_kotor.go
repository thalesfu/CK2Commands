package travunia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科托尔KotorBarony struct {
	feud.BaseBarony
}

var BKotor科托尔 feud.Barony = &科托尔KotorBarony{}

func init() {
	f := BKotor科托尔.(*科托尔KotorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kotor",
		TitleName: "科托尔",
		TitleCode: "b_kotor",
	}
}
