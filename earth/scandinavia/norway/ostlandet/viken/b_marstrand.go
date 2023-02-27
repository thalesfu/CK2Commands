package viken

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马斯特兰德MarstrandBarony struct {
	feud.BaseBarony
}

var BMarstrand马斯特兰德 feud.Barony = &马斯特兰德MarstrandBarony{}

func init() {
    f := BMarstrand马斯特兰德.(*马斯特兰德MarstrandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marstrand",
		TitleName: "马斯特兰德",
		TitleCode: "b_marstrand",
	}
}
