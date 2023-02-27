package gnieznienskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克齐尼亚KcyniaBarony struct {
	feud.BaseBarony
}

var BKcynia克齐尼亚 feud.Barony = &克齐尼亚KcyniaBarony{}

func init() {
    f := BKcynia克齐尼亚.(*克齐尼亚KcyniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kcynia",
		TitleName: "克齐尼亚",
		TitleCode: "b_kcynia",
	}
}
