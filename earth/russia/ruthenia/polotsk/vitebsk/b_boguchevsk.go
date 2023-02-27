package vitebsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博古舍夫斯克BoguchevskBarony struct {
	feud.BaseBarony
}

var BBoguchevsk博古舍夫斯克 feud.Barony = &博古舍夫斯克BoguchevskBarony{}

func init() {
    f := BBoguchevsk博古舍夫斯克.(*博古舍夫斯克BoguchevskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boguchevsk",
		TitleName: "博古舍夫斯克",
		TitleCode: "b_boguchevsk",
	}
}
