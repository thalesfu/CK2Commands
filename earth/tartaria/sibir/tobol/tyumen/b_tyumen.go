package tyumen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 秋明TyumenBarony struct {
	feud.BaseBarony
}

var BTyumen秋明 feud.Barony = &秋明TyumenBarony{}

func init() {
    f := BTyumen秋明.(*秋明TyumenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tyumen",
		TitleName: "秋明",
		TitleCode: "b_tyumen",
	}
}
