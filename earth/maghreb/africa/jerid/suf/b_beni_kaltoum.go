package suf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尼卡尔图姆Beni_kaltoumBarony struct {
	feud.BaseBarony
}

var BBeni_kaltoum贝尼卡尔图姆 feud.Barony = &贝尼卡尔图姆Beni_kaltoumBarony{}

func init() {
    f := BBeni_kaltoum贝尼卡尔图姆.(*贝尼卡尔图姆Beni_kaltoumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beni_kaltoum",
		TitleName: "贝尼卡尔图姆",
		TitleCode: "b_beni_kaltoum",
	}
}
