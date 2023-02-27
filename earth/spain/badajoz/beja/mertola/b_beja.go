package mertola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝雅BejaBarony struct {
	feud.BaseBarony
}

var BBeja贝雅 feud.Barony = &贝雅BejaBarony{}

func init() {
    f := BBeja贝雅.(*贝雅BejaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beja",
		TitleName: "贝雅",
		TitleCode: "b_beja",
	}
}
