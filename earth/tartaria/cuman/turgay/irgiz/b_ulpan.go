package irgiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔潘UlpanBarony struct {
	feud.BaseBarony
}

var BUlpan乌尔潘 feud.Barony = &乌尔潘UlpanBarony{}

func init() {
	f := BUlpan乌尔潘.(*乌尔潘UlpanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulpan",
		TitleName: "乌尔潘",
		TitleCode: "b_ulpan",
	}
}
