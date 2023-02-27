package kalaus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扬库尔YankulBarony struct {
	feud.BaseBarony
}

var BYankul扬库尔 feud.Barony = &扬库尔YankulBarony{}

func init() {
    f := BYankul扬库尔.(*扬库尔YankulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yankul",
		TitleName: "扬库尔",
		TitleCode: "b_yankul",
	}
}
