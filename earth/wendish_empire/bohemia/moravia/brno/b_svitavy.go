package brno

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯维塔维SvitavyBarony struct {
	feud.BaseBarony
}

var BSvitavy斯维塔维 feud.Barony = &斯维塔维SvitavyBarony{}

func init() {
    f := BSvitavy斯维塔维.(*斯维塔维SvitavyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "svitavy",
		TitleName: "斯维塔维",
		TitleCode: "b_svitavy",
	}
}
