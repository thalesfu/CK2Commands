package lubelska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌任杜夫UrzedowBarony struct {
	feud.BaseBarony
}

var BUrzedow乌任杜夫 feud.Barony = &乌任杜夫UrzedowBarony{}

func init() {
    f := BUrzedow乌任杜夫.(*乌任杜夫UrzedowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urzedow",
		TitleName: "乌任杜夫",
		TitleCode: "b_urzedow",
	}
}
