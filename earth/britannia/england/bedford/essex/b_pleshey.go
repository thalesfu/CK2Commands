package essex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布莱希PlesheyBarony struct {
	feud.BaseBarony
}

var BPleshey布莱希 feud.Barony = &布莱希PlesheyBarony{}

func init() {
    f := BPleshey布莱希.(*布莱希PlesheyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pleshey",
		TitleName: "布莱希",
		TitleCode: "b_pleshey",
	}
}
