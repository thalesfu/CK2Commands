package sambhal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆掣罗蕴BachhraonBarony struct {
	feud.BaseBarony
}

var BBachhraon婆掣罗蕴 feud.Barony = &婆掣罗蕴BachhraonBarony{}

func init() {
	f := BBachhraon婆掣罗蕴.(*婆掣罗蕴BachhraonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bachhraon",
		TitleName: "婆掣罗蕴",
		TitleCode: "b_bachhraon",
	}
}
