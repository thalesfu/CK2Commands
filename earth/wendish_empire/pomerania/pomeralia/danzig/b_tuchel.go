package danzig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图赫尔TuchelBarony struct {
	feud.BaseBarony
}

var BTuchel图赫尔 feud.Barony = &图赫尔TuchelBarony{}

func init() {
    f := BTuchel图赫尔.(*图赫尔TuchelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tuchel",
		TitleName: "图赫尔",
		TitleCode: "b_tuchel",
	}
}
