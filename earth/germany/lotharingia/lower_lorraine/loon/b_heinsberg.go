package loon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海因斯贝格HeinsbergBarony struct {
	feud.BaseBarony
}

var BHeinsberg海因斯贝格 feud.Barony = &海因斯贝格HeinsbergBarony{}

func init() {
    f := BHeinsberg海因斯贝格.(*海因斯贝格HeinsbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "heinsberg",
		TitleName: "海因斯贝格",
		TitleCode: "b_heinsberg",
	}
}
