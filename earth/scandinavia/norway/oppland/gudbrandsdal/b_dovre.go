package gudbrandsdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多夫勒DovreBarony struct {
	feud.BaseBarony
}

var BDovre多夫勒 feud.Barony = &多夫勒DovreBarony{}

func init() {
    f := BDovre多夫勒.(*多夫勒DovreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dovre",
		TitleName: "多夫勒",
		TitleCode: "b_dovre",
	}
}
