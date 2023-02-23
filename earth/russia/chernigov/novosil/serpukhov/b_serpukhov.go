package serpukhov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢尔普霍夫SerpukhovBarony struct {
	feud.BaseBarony
}

var BSerpukhov谢尔普霍夫 feud.Barony = &谢尔普霍夫SerpukhovBarony{}

func init() {
	f := BSerpukhov谢尔普霍夫.(*谢尔普霍夫SerpukhovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "serpukhov",
		TitleName: "谢尔普霍夫",
		TitleCode: "b_serpukhov",
	}
}
