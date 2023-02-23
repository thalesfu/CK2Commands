package ramagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 难提拔檀那NandivardhanaBarony struct {
	feud.BaseBarony
}

var BNandivardhana难提拔檀那 feud.Barony = &难提拔檀那NandivardhanaBarony{}

func init() {
	f := BNandivardhana难提拔檀那.(*难提拔檀那NandivardhanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nandivardhana",
		TitleName: "难提拔檀那",
		TitleCode: "b_nandivardhana",
	}
}
