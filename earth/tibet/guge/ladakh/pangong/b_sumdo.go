package pangong

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏姆多SumdoBarony struct {
	feud.BaseBarony
}

var BSumdo苏姆多 feud.Barony = &苏姆多SumdoBarony{}

func init() {
    f := BSumdo苏姆多.(*苏姆多SumdoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sumdo",
		TitleName: "苏姆多",
		TitleCode: "b_sumdo",
	}
}
