package kherla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝玛扬泽PemayangtseBarony struct {
	feud.BaseBarony
}

var BPemayangtse贝玛扬泽 feud.Barony = &贝玛扬泽PemayangtseBarony{}

func init() {
    f := BPemayangtse贝玛扬泽.(*贝玛扬泽PemayangtseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pemayangtse",
		TitleName: "贝玛扬泽",
		TitleCode: "b_pemayangtse",
	}
}
