package yegorlyk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔肯_沙哈尔ErkenshakharBarony struct {
	feud.BaseBarony
}

var BErkenshakhar埃尔肯_沙哈尔 feud.Barony = &埃尔肯_沙哈尔ErkenshakharBarony{}

func init() {
    f := BErkenshakhar埃尔肯_沙哈尔.(*埃尔肯_沙哈尔ErkenshakharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erkenshakhar",
		TitleName: "埃尔肯_沙哈尔",
		TitleCode: "b_erkenshakhar",
	}
}
