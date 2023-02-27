package beshbaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 将军庙JiangjunmiaoBarony struct {
	feud.BaseBarony
}

var BJiangjunmiao将军庙 feud.Barony = &将军庙JiangjunmiaoBarony{}

func init() {
    f := BJiangjunmiao将军庙.(*将军庙JiangjunmiaoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jiangjunmiao",
		TitleName: "将军庙",
		TitleCode: "b_jiangjunmiao",
	}
}
