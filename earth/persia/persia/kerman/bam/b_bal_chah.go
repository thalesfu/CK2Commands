package bam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴拉察Bal_chahBarony struct {
	feud.BaseBarony
}

var BBal_chah巴拉察 feud.Barony = &巴拉察Bal_chahBarony{}

func init() {
    f := BBal_chah巴拉察.(*巴拉察Bal_chahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bal_chah",
		TitleName: "巴拉察",
		TitleCode: "b_bal_chah",
	}
}
