package orekhovo_zouievo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥列霍沃_祖耶沃Orekhovo_zouievoBarony struct {
	feud.BaseBarony
}

var BOrekhovo_zouievo奥列霍沃_祖耶沃 feud.Barony = &奥列霍沃_祖耶沃Orekhovo_zouievoBarony{}

func init() {
    f := BOrekhovo_zouievo奥列霍沃_祖耶沃.(*奥列霍沃_祖耶沃Orekhovo_zouievoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orekhovo_zouievo",
		TitleName: "奥列霍沃_祖耶沃",
		TitleCode: "b_orekhovo-zouievo",
	}
}
