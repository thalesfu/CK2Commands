package mesopotamia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪亚巴克尔TigranakertBarony struct {
	feud.BaseBarony
}

var BTigranakert迪亚巴克尔 feud.Barony = &迪亚巴克尔TigranakertBarony{}

func init() {
    f := BTigranakert迪亚巴克尔.(*迪亚巴克尔TigranakertBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tigranakert",
		TitleName: "迪亚巴克尔",
		TitleCode: "b_tigranakert",
	}
}
