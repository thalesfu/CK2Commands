package zanjan_abhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古兹卢切QuzluchehBarony struct {
	feud.BaseBarony
}

var BQuzlucheh古兹卢切 feud.Barony = &古兹卢切QuzluchehBarony{}

func init() {
    f := BQuzlucheh古兹卢切.(*古兹卢切QuzluchehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "quzlucheh",
		TitleName: "古兹卢切",
		TitleCode: "b_quzlucheh",
	}
}
