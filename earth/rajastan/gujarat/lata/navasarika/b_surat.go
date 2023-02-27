package navasarika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏剌吒SuratBarony struct {
	feud.BaseBarony
}

var BSurat苏剌吒 feud.Barony = &苏剌吒SuratBarony{}

func init() {
    f := BSurat苏剌吒.(*苏剌吒SuratBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "surat",
		TitleName: "苏剌吒",
		TitleCode: "b_surat",
	}
}
