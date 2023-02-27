package neuchatel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦朗然ValanginBarony struct {
	feud.BaseBarony
}

var BValangin瓦朗然 feud.Barony = &瓦朗然ValanginBarony{}

func init() {
    f := BValangin瓦朗然.(*瓦朗然ValanginBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valangin",
		TitleName: "瓦朗然",
		TitleCode: "b_valangin",
	}
}
