package plzen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉德鲁比KladrubyBarony struct {
	feud.BaseBarony
}

var BKladruby克拉德鲁比 feud.Barony = &克拉德鲁比KladrubyBarony{}

func init() {
    f := BKladruby克拉德鲁比.(*克拉德鲁比KladrubyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kladruby",
		TitleName: "克拉德鲁比",
		TitleCode: "b_kladruby",
	}
}
