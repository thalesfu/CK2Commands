package hanyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳马NaamaBarony struct {
	feud.BaseBarony
}

var BNaama纳马 feud.Barony = &纳马NaamaBarony{}

func init() {
    f := BNaama纳马.(*纳马NaamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naama",
		TitleName: "纳马",
		TitleCode: "b_naama",
	}
}
