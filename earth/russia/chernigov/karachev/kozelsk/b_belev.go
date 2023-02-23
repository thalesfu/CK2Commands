package kozelsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别列夫BelevBarony struct {
	feud.BaseBarony
}

var BBelev别列夫 feud.Barony = &别列夫BelevBarony{}

func init() {
	f := BBelev别列夫.(*别列夫BelevBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belev",
		TitleName: "别列夫",
		TitleCode: "b_belev",
	}
}
