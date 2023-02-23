package arques

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑞米耶日JumiegesBarony struct {
	feud.BaseBarony
}

var BJumieges瑞米耶日 feud.Barony = &瑞米耶日JumiegesBarony{}

func init() {
	f := BJumieges瑞米耶日.(*瑞米耶日JumiegesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jumieges",
		TitleName: "瑞米耶日",
		TitleCode: "b_jumieges",
	}
}
