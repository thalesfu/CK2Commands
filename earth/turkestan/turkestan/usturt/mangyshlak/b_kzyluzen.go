package mangyshlak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克孜勒乌赞KzyluzenBarony struct {
	feud.BaseBarony
}

var BKzyluzen克孜勒乌赞 feud.Barony = &克孜勒乌赞KzyluzenBarony{}

func init() {
	f := BKzyluzen克孜勒乌赞.(*克孜勒乌赞KzyluzenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kzyluzen",
		TitleName: "克孜勒乌赞",
		TitleCode: "b_kzyluzen",
	}
}
