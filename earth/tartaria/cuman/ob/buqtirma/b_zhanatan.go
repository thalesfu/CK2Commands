package buqtirma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎纳坦ZhanatanBarony struct {
	feud.BaseBarony
}

var BZhanatan扎纳坦 feud.Barony = &扎纳坦ZhanatanBarony{}

func init() {
    f := BZhanatan扎纳坦.(*扎纳坦ZhanatanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhanatan",
		TitleName: "扎纳坦",
		TitleCode: "b_zhanatan",
	}
}
