package fars

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈夫尔KhafrBarony struct {
	feud.BaseBarony
}

var BKhafr哈夫尔 feud.Barony = &哈夫尔KhafrBarony{}

func init() {
	f := BKhafr哈夫尔.(*哈夫尔KhafrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khafr",
		TitleName: "哈夫尔",
		TitleCode: "b_khafr",
	}
}
