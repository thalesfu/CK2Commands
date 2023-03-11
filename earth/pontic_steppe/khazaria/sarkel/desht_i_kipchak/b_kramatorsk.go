package desht_i_kipchak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉马托尔斯克KramatorskBarony struct {
	feud.BaseBarony
}

var BKramatorsk克拉马托尔斯克 feud.Barony = &克拉马托尔斯克KramatorskBarony{}

func init() {
    f := BKramatorsk克拉马托尔斯克.(*克拉马托尔斯克KramatorskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kramatorsk",
		TitleName: "克拉马托尔斯克",
		TitleCode: "b_kramatorsk",
	}
}
