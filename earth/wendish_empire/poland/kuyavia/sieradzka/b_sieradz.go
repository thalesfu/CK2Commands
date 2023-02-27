package sieradzka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢拉兹SieradzBarony struct {
	feud.BaseBarony
}

var BSieradz谢拉兹 feud.Barony = &谢拉兹SieradzBarony{}

func init() {
    f := BSieradz谢拉兹.(*谢拉兹SieradzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sieradz",
		TitleName: "谢拉兹",
		TitleCode: "b_sieradz",
	}
}
