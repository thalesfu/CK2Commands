package delingha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柯鲁柯KelukeBarony struct {
	feud.BaseBarony
}

var BKeluke柯鲁柯 feud.Barony = &柯鲁柯KelukeBarony{}

func init() {
	f := BKeluke柯鲁柯.(*柯鲁柯KelukeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "keluke",
		TitleName: "柯鲁柯",
		TitleCode: "b_keluke",
	}
}
