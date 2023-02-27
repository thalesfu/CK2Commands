package orleans

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 让维尔JanvilleBarony struct {
	feud.BaseBarony
}

var BJanville让维尔 feud.Barony = &让维尔JanvilleBarony{}

func init() {
    f := BJanville让维尔.(*让维尔JanvilleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "janville",
		TitleName: "让维尔",
		TitleCode: "b_janville",
	}
}
