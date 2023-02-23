package katun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔卢什卡KarlushkaBarony struct {
	feud.BaseBarony
}

var BKarlushka卡尔卢什卡 feud.Barony = &卡尔卢什卡KarlushkaBarony{}

func init() {
	f := BKarlushka卡尔卢什卡.(*卡尔卢什卡KarlushkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karlushka",
		TitleName: "卡尔卢什卡",
		TitleCode: "b_karlushka",
	}
}
