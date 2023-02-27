package massawa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科哈依托QohaitoBarony struct {
	feud.BaseBarony
}

var BQohaito科哈依托 feud.Barony = &科哈依托QohaitoBarony{}

func init() {
    f := BQohaito科哈依托.(*科哈依托QohaitoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qohaito",
		TitleName: "科哈依托",
		TitleCode: "b_qohaito",
	}
}
