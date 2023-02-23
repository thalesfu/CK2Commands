package sennar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦德拉瓦赫WadrawahBarony struct {
	feud.BaseBarony
}

var BWadrawah瓦德拉瓦赫 feud.Barony = &瓦德拉瓦赫WadrawahBarony{}

func init() {
	f := BWadrawah瓦德拉瓦赫.(*瓦德拉瓦赫WadrawahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wadrawah",
		TitleName: "瓦德拉瓦赫",
		TitleCode: "b_wadrawah",
	}
}
