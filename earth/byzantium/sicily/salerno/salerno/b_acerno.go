package salerno

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿切尔诺AcernoBarony struct {
	feud.BaseBarony
}

var BAcerno阿切尔诺 feud.Barony = &阿切尔诺AcernoBarony{}

func init() {
    f := BAcerno阿切尔诺.(*阿切尔诺AcernoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "acerno",
		TitleName: "阿切尔诺",
		TitleCode: "b_acerno",
	}
}
