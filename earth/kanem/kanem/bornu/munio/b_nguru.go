package munio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩古鲁NguruBarony struct {
	feud.BaseBarony
}

var BNguru恩古鲁 feud.Barony = &恩古鲁NguruBarony{}

func init() {
	f := BNguru恩古鲁.(*恩古鲁NguruBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nguru",
		TitleName: "恩古鲁",
		TitleCode: "b_nguru",
	}
}
