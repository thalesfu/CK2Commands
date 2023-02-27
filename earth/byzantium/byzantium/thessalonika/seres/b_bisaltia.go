package seres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比萨尔提亚BisaltiaBarony struct {
	feud.BaseBarony
}

var BBisaltia比萨尔提亚 feud.Barony = &比萨尔提亚BisaltiaBarony{}

func init() {
    f := BBisaltia比萨尔提亚.(*比萨尔提亚BisaltiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bisaltia",
		TitleName: "比萨尔提亚",
		TitleCode: "b_bisaltia",
	}
}
