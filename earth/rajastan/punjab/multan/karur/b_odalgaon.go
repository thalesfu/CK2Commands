package karur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌陀罗伽罗摩OdalgaonBarony struct {
	feud.BaseBarony
}

var BOdalgaon乌陀罗伽罗摩 feud.Barony = &乌陀罗伽罗摩OdalgaonBarony{}

func init() {
    f := BOdalgaon乌陀罗伽罗摩.(*乌陀罗伽罗摩OdalgaonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "odalgaon",
		TitleName: "乌陀罗伽罗摩",
		TitleCode: "b_odalgaon",
	}
}
