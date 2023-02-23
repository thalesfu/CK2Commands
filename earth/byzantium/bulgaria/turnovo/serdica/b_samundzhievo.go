package serdica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨蒙吉耶沃SamundzhievoBarony struct {
	feud.BaseBarony
}

var BSamundzhievo萨蒙吉耶沃 feud.Barony = &萨蒙吉耶沃SamundzhievoBarony{}

func init() {
	f := BSamundzhievo萨蒙吉耶沃.(*萨蒙吉耶沃SamundzhievoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samundzhievo",
		TitleName: "萨蒙吉耶沃",
		TitleCode: "b_samundzhievo",
	}
}
