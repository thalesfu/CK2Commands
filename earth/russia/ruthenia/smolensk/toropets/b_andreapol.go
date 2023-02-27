package toropets

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安德烈亚波尔AndreapolBarony struct {
	feud.BaseBarony
}

var BAndreapol安德烈亚波尔 feud.Barony = &安德烈亚波尔AndreapolBarony{}

func init() {
    f := BAndreapol安德烈亚波尔.(*安德烈亚波尔AndreapolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "andreapol",
		TitleName: "安德烈亚波尔",
		TitleCode: "b_andreapol",
	}
}
