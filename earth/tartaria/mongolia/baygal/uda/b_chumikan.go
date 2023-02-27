package uda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丘米坎ChumikanBarony struct {
	feud.BaseBarony
}

var BChumikan丘米坎 feud.Barony = &丘米坎ChumikanBarony{}

func init() {
    f := BChumikan丘米坎.(*丘米坎ChumikanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chumikan",
		TitleName: "丘米坎",
		TitleCode: "b_chumikan",
	}
}
