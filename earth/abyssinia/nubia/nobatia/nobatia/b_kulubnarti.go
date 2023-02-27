package nobatia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库勒卜KulubnartiBarony struct {
	feud.BaseBarony
}

var BKulubnarti库勒卜 feud.Barony = &库勒卜KulubnartiBarony{}

func init() {
    f := BKulubnarti库勒卜.(*库勒卜KulubnartiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kulubnarti",
		TitleName: "库勒卜",
		TitleCode: "b_kulubnarti",
	}
}
