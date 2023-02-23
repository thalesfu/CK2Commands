package adrianopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈德良波利斯AdrianopolisBarony struct {
	feud.BaseBarony
}

var BAdrianopolis哈德良波利斯 feud.Barony = &哈德良波利斯AdrianopolisBarony{}

func init() {
	f := BAdrianopolis哈德良波利斯.(*哈德良波利斯AdrianopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adrianopolis",
		TitleName: "哈德良波利斯",
		TitleCode: "b_adrianopolis",
	}
}
