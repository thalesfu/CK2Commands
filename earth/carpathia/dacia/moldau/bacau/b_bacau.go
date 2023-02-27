package bacau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴克乌BacauBarony struct {
	feud.BaseBarony
}

var BBacau巴克乌 feud.Barony = &巴克乌BacauBarony{}

func init() {
    f := BBacau巴克乌.(*巴克乌BacauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bacau",
		TitleName: "巴克乌",
		TitleCode: "b_bacau",
	}
}
