package gent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥德纳尔德OudenaardeBarony struct {
	feud.BaseBarony
}

var BOudenaarde奥德纳尔德 feud.Barony = &奥德纳尔德OudenaardeBarony{}

func init() {
    f := BOudenaarde奥德纳尔德.(*奥德纳尔德OudenaardeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oudenaarde",
		TitleName: "奥德纳尔德",
		TitleCode: "b_oudenaarde",
	}
}
