package ratanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 势伽罗摩ShegaonBarony struct {
	feud.BaseBarony
}

var BShegaon势伽罗摩 feud.Barony = &势伽罗摩ShegaonBarony{}

func init() {
	f := BShegaon势伽罗摩.(*势伽罗摩ShegaonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shegaon",
		TitleName: "势伽罗摩",
		TitleCode: "b_shegaon",
	}
}
