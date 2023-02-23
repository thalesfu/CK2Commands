package almansa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 考德特CaudeteBarony struct {
	feud.BaseBarony
}

var BCaudete考德特 feud.Barony = &考德特CaudeteBarony{}

func init() {
	f := BCaudete考德特.(*考德特CaudeteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caudete",
		TitleName: "考德特",
		TitleCode: "b_caudete",
	}
}
