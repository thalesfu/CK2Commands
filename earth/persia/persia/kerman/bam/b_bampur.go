package bam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班布尔BampurBarony struct {
	feud.BaseBarony
}

var BBampur班布尔 feud.Barony = &班布尔BampurBarony{}

func init() {
    f := BBampur班布尔.(*班布尔BampurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bampur",
		TitleName: "班布尔",
		TitleCode: "b_bampur",
	}
}
