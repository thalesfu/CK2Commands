package bhakkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩吉阿MejiaBarony struct {
	feud.BaseBarony
}

var BMejia摩吉阿 feud.Barony = &摩吉阿MejiaBarony{}

func init() {
    f := BMejia摩吉阿.(*摩吉阿MejiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mejia",
		TitleName: "摩吉阿",
		TitleCode: "b_mejia",
	}
}
