package travunia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣斯特凡Sv_stefanBarony struct {
	feud.BaseBarony
}

var BSv_stefan圣斯特凡 feud.Barony = &圣斯特凡Sv_stefanBarony{}

func init() {
    f := BSv_stefan圣斯特凡.(*圣斯特凡Sv_stefanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sv_stefan",
		TitleName: "圣斯特凡",
		TitleCode: "b_sv_stefan",
	}
}
