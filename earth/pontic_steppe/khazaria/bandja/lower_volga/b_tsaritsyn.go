package lower_volga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 察里津TsaritsynBarony struct {
	feud.BaseBarony
}

var BTsaritsyn察里津 feud.Barony = &察里津TsaritsynBarony{}

func init() {
    f := BTsaritsyn察里津.(*察里津TsaritsynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsaritsyn",
		TitleName: "察里津",
		TitleCode: "b_tsaritsyn",
	}
}
