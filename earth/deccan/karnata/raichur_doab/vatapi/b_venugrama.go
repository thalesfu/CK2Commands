package vatapi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鞞纽伽罗摩VenugramaBarony struct {
	feud.BaseBarony
}

var BVenugrama鞞纽伽罗摩 feud.Barony = &鞞纽伽罗摩VenugramaBarony{}

func init() {
    f := BVenugrama鞞纽伽罗摩.(*鞞纽伽罗摩VenugramaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "venugrama",
		TitleName: "鞞纽伽罗摩",
		TitleCode: "b_venugrama",
	}
}
