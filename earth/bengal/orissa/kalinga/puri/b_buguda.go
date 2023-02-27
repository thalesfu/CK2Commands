package puri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菩瞿陀BugudaBarony struct {
	feud.BaseBarony
}

var BBuguda菩瞿陀 feud.Barony = &菩瞿陀BugudaBarony{}

func init() {
    f := BBuguda菩瞿陀.(*菩瞿陀BugudaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buguda",
		TitleName: "菩瞿陀",
		TitleCode: "b_buguda",
	}
}
