package kromy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢夫斯克SevskBarony struct {
	feud.BaseBarony
}

var BSevsk谢夫斯克 feud.Barony = &谢夫斯克SevskBarony{}

func init() {
    f := BSevsk谢夫斯克.(*谢夫斯克SevskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sevsk",
		TitleName: "谢夫斯克",
		TitleCode: "b_sevsk",
	}
}
