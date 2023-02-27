package thrake

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 撒米德索斯SalmydessusBarony struct {
	feud.BaseBarony
}

var BSalmydessus撒米德索斯 feud.Barony = &撒米德索斯SalmydessusBarony{}

func init() {
    f := BSalmydessus撒米德索斯.(*撒米德索斯SalmydessusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salmydessus",
		TitleName: "撒米德索斯",
		TitleCode: "b_salmydessus",
	}
}
