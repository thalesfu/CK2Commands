package yeruslan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊玛什ImashBarony struct {
	feud.BaseBarony
}

var BImash伊玛什 feud.Barony = &伊玛什ImashBarony{}

func init() {
    f := BImash伊玛什.(*伊玛什ImashBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "imash",
		TitleName: "伊玛什",
		TitleCode: "b_imash",
	}
}
