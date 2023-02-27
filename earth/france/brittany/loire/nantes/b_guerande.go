package nantes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖朗德GuerandeBarony struct {
	feud.BaseBarony
}

var BGuerande盖朗德 feud.Barony = &盖朗德GuerandeBarony{}

func init() {
    f := BGuerande盖朗德.(*盖朗德GuerandeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guerande",
		TitleName: "盖朗德",
		TitleCode: "b_guerande",
	}
}
