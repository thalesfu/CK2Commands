package tis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕姆马处汗Pozm_machchanBarony struct {
	feud.BaseBarony
}

var BPozm_machchan帕姆马处汗 feud.Barony = &帕姆马处汗Pozm_machchanBarony{}

func init() {
    f := BPozm_machchan帕姆马处汗.(*帕姆马处汗Pozm_machchanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pozm_machchan",
		TitleName: "帕姆马处汗",
		TitleCode: "b_pozm_machchan",
	}
}
