package bar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茹安维尔JoinvilleBarony struct {
	feud.BaseBarony
}

var BJoinville茹安维尔 feud.Barony = &茹安维尔JoinvilleBarony{}

func init() {
    f := BJoinville茹安维尔.(*茹安维尔JoinvilleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "joinville",
		TitleName: "茹安维尔",
		TitleCode: "b_joinville",
	}
}
