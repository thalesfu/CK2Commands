package orekhovo-zouievo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥列霍沃_祖耶沃Orekhovo-zouievoBarony struct {
	feud.BaseBarony
}

var BOrekhovo-zouievo奥列霍沃_祖耶沃 feud.Barony = &奥列霍沃_祖耶沃Orekhovo-zouievoBarony{}

func init() {
    f := BOrekhovo-zouievo奥列霍沃_祖耶沃.(*奥列霍沃_祖耶沃Orekhovo-zouievoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orekhovo-zouievo",
		TitleName: "奥列霍沃_祖耶沃",
		TitleCode: "b_orekhovo-zouievo",
	}
}
