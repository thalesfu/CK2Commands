package napata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔提KortiBarony struct {
	feud.BaseBarony
}

var BKorti库尔提 feud.Barony = &库尔提KortiBarony{}

func init() {
    f := BKorti库尔提.(*库尔提KortiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korti",
		TitleName: "库尔提",
		TitleCode: "b_korti",
	}
}
