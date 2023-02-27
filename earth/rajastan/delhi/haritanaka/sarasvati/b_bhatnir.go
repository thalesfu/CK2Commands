package sarasvati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆提尼BhatnirBarony struct {
	feud.BaseBarony
}

var BBhatnir婆提尼 feud.Barony = &婆提尼BhatnirBarony{}

func init() {
    f := BBhatnir婆提尼.(*婆提尼BhatnirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhatnir",
		TitleName: "婆提尼",
		TitleCode: "b_bhatnir",
	}
}
