package biskra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴特纳BatnaBarony struct {
	feud.BaseBarony
}

var BBatna巴特纳 feud.Barony = &巴特纳BatnaBarony{}

func init() {
    f := BBatna巴特纳.(*巴特纳BatnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "batna",
		TitleName: "巴特纳",
		TitleCode: "b_batna",
	}
}
