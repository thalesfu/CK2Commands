package thrake

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈里奥波利斯ChariopolisBarony struct {
	feud.BaseBarony
}

var BChariopolis哈里奥波利斯 feud.Barony = &哈里奥波利斯ChariopolisBarony{}

func init() {
    f := BChariopolis哈里奥波利斯.(*哈里奥波利斯ChariopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chariopolis",
		TitleName: "哈里奥波利斯",
		TitleCode: "b_chariopolis",
	}
}
