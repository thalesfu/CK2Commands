package novogrodek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 新格鲁多克NovogrodekBarony struct {
	feud.BaseBarony
}

var BNovogrodek新格鲁多克 feud.Barony = &新格鲁多克NovogrodekBarony{}

func init() {
	f := BNovogrodek新格鲁多克.(*新格鲁多克NovogrodekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "novogrodek",
		TitleName: "新格鲁多克",
		TitleCode: "b_novogrodek",
	}
}
