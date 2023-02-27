package kara_bogaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝克达什BekdashBarony struct {
	feud.BaseBarony
}

var BBekdash贝克达什 feud.Barony = &贝克达什BekdashBarony{}

func init() {
    f := BBekdash贝克达什.(*贝克达什BekdashBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bekdash",
		TitleName: "贝克达什",
		TitleCode: "b_bekdash",
	}
}
