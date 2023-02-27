package nagauda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩伽罗那MakranaBarony struct {
	feud.BaseBarony
}

var BMakrana摩伽罗那 feud.Barony = &摩伽罗那MakranaBarony{}

func init() {
    f := BMakrana摩伽罗那.(*摩伽罗那MakranaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "makrana",
		TitleName: "摩伽罗那",
		TitleCode: "b_makrana",
	}
}
