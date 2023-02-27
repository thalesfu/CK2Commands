package st_gallen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利希滕施泰格LichtensteigBarony struct {
	feud.BaseBarony
}

var BLichtensteig利希滕施泰格 feud.Barony = &利希滕施泰格LichtensteigBarony{}

func init() {
    f := BLichtensteig利希滕施泰格.(*利希滕施泰格LichtensteigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lichtensteig",
		TitleName: "利希滕施泰格",
		TitleCode: "b_lichtensteig",
	}
}
