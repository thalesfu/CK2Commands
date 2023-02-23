package burhanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 折罗伽罗摩JargaonBarony struct {
	feud.BaseBarony
}

var BJargaon折罗伽罗摩 feud.Barony = &折罗伽罗摩JargaonBarony{}

func init() {
	f := BJargaon折罗伽罗摩.(*折罗伽罗摩JargaonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jargaon",
		TitleName: "折罗伽罗摩",
		TitleCode: "b_jargaon",
	}
}
