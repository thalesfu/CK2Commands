package usora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科托尔Kotor_donjikrajiBarony struct {
	feud.BaseBarony
}

var BKotor_donjikraji科托尔 feud.Barony = &科托尔Kotor_donjikrajiBarony{}

func init() {
    f := BKotor_donjikraji科托尔.(*科托尔Kotor_donjikrajiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kotor_donjikraji",
		TitleName: "科托尔",
		TitleCode: "b_kotor_donjikraji",
	}
}
