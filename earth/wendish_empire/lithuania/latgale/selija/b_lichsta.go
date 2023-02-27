package selija

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利奇斯塔LichstaBarony struct {
	feud.BaseBarony
}

var BLichsta利奇斯塔 feud.Barony = &利奇斯塔LichstaBarony{}

func init() {
    f := BLichsta利奇斯塔.(*利奇斯塔LichstaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lichsta",
		TitleName: "利奇斯塔",
		TitleCode: "b_lichsta",
	}
}
