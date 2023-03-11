package bryansk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特鲁别茨克TrubetskBarony struct {
	feud.BaseBarony
}

var BTrubetsk特鲁别茨克 feud.Barony = &特鲁别茨克TrubetskBarony{}

func init() {
    f := BTrubetsk特鲁别茨克.(*特鲁别茨克TrubetskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trubetsk",
		TitleName: "特鲁别茨克",
		TitleCode: "b_trubetsk",
	}
}
