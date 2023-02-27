package lubusz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏伦恰SuleciaBarony struct {
	feud.BaseBarony
}

var BSulecia苏伦恰 feud.Barony = &苏伦恰SuleciaBarony{}

func init() {
    f := BSulecia苏伦恰.(*苏伦恰SuleciaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sulecia",
		TitleName: "苏伦恰",
		TitleCode: "b_sulecia",
	}
}
