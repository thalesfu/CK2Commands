package bornu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赞坦ZamtamBarony struct {
	feud.BaseBarony
}

var BZamtam赞坦 feud.Barony = &赞坦ZamtamBarony{}

func init() {
    f := BZamtam赞坦.(*赞坦ZamtamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zamtam",
		TitleName: "赞坦",
		TitleCode: "b_zamtam",
	}
}
