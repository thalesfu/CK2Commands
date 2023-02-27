package gorlitz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊卢施蒂亚贝格YlustiabergBarony struct {
	feud.BaseBarony
}

var BYlustiaberg伊卢施蒂亚贝格 feud.Barony = &伊卢施蒂亚贝格YlustiabergBarony{}

func init() {
    f := BYlustiaberg伊卢施蒂亚贝格.(*伊卢施蒂亚贝格YlustiabergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ylustiaberg",
		TitleName: "伊卢施蒂亚贝格",
		TitleCode: "b_ylustiaberg",
	}
}
