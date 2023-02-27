package kerman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努格NoughBarony struct {
	feud.BaseBarony
}

var BNough努格 feud.Barony = &努格NoughBarony{}

func init() {
    f := BNough努格.(*努格NoughBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nough",
		TitleName: "努格",
		TitleCode: "b_nough",
	}
}
