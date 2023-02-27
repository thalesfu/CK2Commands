package syrmia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博诺尼亚BononiaBarony struct {
	feud.BaseBarony
}

var BBononia博诺尼亚 feud.Barony = &博诺尼亚BononiaBarony{}

func init() {
    f := BBononia博诺尼亚.(*博诺尼亚BononiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bononia",
		TitleName: "博诺尼亚",
		TitleCode: "b_bononia",
	}
}
