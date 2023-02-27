package khlynov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡巴奇希KabachischiBarony struct {
	feud.BaseBarony
}

var BKabachischi卡巴奇希 feud.Barony = &卡巴奇希KabachischiBarony{}

func init() {
    f := BKabachischi卡巴奇希.(*卡巴奇希KabachischiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kabachischi",
		TitleName: "卡巴奇希",
		TitleCode: "b_kabachischi",
	}
}
