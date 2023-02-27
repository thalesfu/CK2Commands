package atheniai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡律季KarydiBarony struct {
	feud.BaseBarony
}

var BKarydi卡律季 feud.Barony = &卡律季KarydiBarony{}

func init() {
    f := BKarydi卡律季.(*卡律季KarydiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karydi",
		TitleName: "卡律季",
		TitleCode: "b_karydi",
	}
}
