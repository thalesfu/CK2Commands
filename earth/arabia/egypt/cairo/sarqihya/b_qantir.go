package sarqihya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎提尔QantirBarony struct {
	feud.BaseBarony
}

var BQantir坎提尔 feud.Barony = &坎提尔QantirBarony{}

func init() {
	f := BQantir坎提尔.(*坎提尔QantirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qantir",
		TitleName: "坎提尔",
		TitleCode: "b_qantir",
	}
}
