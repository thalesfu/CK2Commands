package ushytsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌希因UshiyinBarony struct {
	feud.BaseBarony
}

var BUshiyin乌希因 feud.Barony = &乌希因UshiyinBarony{}

func init() {
    f := BUshiyin乌希因.(*乌希因UshiyinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ushiyin",
		TitleName: "乌希因",
		TitleCode: "b_ushiyin",
	}
}
