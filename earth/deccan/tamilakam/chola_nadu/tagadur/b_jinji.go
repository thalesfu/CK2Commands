package tagadur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旖耆JinjiBarony struct {
	feud.BaseBarony
}

var BJinji旖耆 feud.Barony = &旖耆JinjiBarony{}

func init() {
    f := BJinji旖耆.(*旖耆JinjiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jinji",
		TitleName: "旖耆",
		TitleCode: "b_jinji",
	}
}
