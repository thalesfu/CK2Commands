package steiermark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大拉明GrobramingBarony struct {
	feud.BaseBarony
}

var BGrobraming大拉明 feud.Barony = &大拉明GrobramingBarony{}

func init() {
    f := BGrobraming大拉明.(*大拉明GrobramingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grobraming",
		TitleName: "大拉明",
		TitleCode: "b_grobraming",
	}
}
