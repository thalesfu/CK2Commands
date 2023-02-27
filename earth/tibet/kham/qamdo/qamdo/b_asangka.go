package qamdo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加桑卡AsangkaBarony struct {
	feud.BaseBarony
}

var BAsangka加桑卡 feud.Barony = &加桑卡AsangkaBarony{}

func init() {
    f := BAsangka加桑卡.(*加桑卡AsangkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asangka",
		TitleName: "加桑卡",
		TitleCode: "b_asangka",
	}
}
