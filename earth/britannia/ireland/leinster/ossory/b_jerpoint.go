package ossory

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰洛因JerpointBarony struct {
	feud.BaseBarony
}

var BJerpoint杰洛因 feud.Barony = &杰洛因JerpointBarony{}

func init() {
    f := BJerpoint杰洛因.(*杰洛因JerpointBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jerpoint",
		TitleName: "杰洛因",
		TitleCode: "b_jerpoint",
	}
}
