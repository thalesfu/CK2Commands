package halland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉内斯AranaesBarony struct {
	feud.BaseBarony
}

var BAranaes阿拉内斯 feud.Barony = &阿拉内斯AranaesBarony{}

func init() {
    f := BAranaes阿拉内斯.(*阿拉内斯AranaesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aranaes",
		TitleName: "阿拉内斯",
		TitleCode: "b_aranaes",
	}
}
