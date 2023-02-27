package lhasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 逻些LhasaBarony struct {
	feud.BaseBarony
}

var BLhasa逻些 feud.Barony = &逻些LhasaBarony{}

func init() {
    f := BLhasa逻些.(*逻些LhasaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lhasa",
		TitleName: "逻些",
		TitleCode: "b_lhasa",
	}
}
