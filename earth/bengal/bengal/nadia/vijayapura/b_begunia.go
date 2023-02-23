package vijayapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗古尼阿BeguniaBarony struct {
	feud.BaseBarony
}

var BBegunia毗古尼阿 feud.Barony = &毗古尼阿BeguniaBarony{}

func init() {
	f := BBegunia毗古尼阿.(*毗古尼阿BeguniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "begunia",
		TitleName: "毗古尼阿",
		TitleCode: "b_begunia",
	}
}
