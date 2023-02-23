package khojand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苦盏KhojandBarony struct {
	feud.BaseBarony
}

var BKhojand苦盏 feud.Barony = &苦盏KhojandBarony{}

func init() {
	f := BKhojand苦盏.(*苦盏KhojandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khojand",
		TitleName: "苦盏",
		TitleCode: "b_khojand",
	}
}
