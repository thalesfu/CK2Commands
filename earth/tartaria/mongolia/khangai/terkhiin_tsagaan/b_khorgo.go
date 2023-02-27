package terkhiin_tsagaan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍尔戈KhorgoBarony struct {
	feud.BaseBarony
}

var BKhorgo霍尔戈 feud.Barony = &霍尔戈KhorgoBarony{}

func init() {
    f := BKhorgo霍尔戈.(*霍尔戈KhorgoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khorgo",
		TitleName: "霍尔戈",
		TitleCode: "b_khorgo",
	}
}
