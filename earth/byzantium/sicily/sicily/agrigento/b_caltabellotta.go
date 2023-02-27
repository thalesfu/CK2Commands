package agrigento

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔塔贝洛塔CaltabellottaBarony struct {
	feud.BaseBarony
}

var BCaltabellotta卡尔塔贝洛塔 feud.Barony = &卡尔塔贝洛塔CaltabellottaBarony{}

func init() {
    f := BCaltabellotta卡尔塔贝洛塔.(*卡尔塔贝洛塔CaltabellottaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caltabellotta",
		TitleName: "卡尔塔贝洛塔",
		TitleCode: "b_caltabellotta",
	}
}
