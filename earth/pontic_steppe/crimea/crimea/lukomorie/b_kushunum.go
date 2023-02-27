package lukomorie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库舒古姆KushunumBarony struct {
	feud.BaseBarony
}

var BKushunum库舒古姆 feud.Barony = &库舒古姆KushunumBarony{}

func init() {
    f := BKushunum库舒古姆.(*库舒古姆KushunumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kushunum",
		TitleName: "库舒古姆",
		TitleCode: "b_kushunum",
	}
}
