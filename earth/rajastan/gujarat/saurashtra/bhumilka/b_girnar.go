package bhumilka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉尔纳尔GirnarBarony struct {
	feud.BaseBarony
}

var BGirnar吉尔纳尔 feud.Barony = &吉尔纳尔GirnarBarony{}

func init() {
    f := BGirnar吉尔纳尔.(*吉尔纳尔GirnarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "girnar",
		TitleName: "吉尔纳尔",
		TitleCode: "b_girnar",
	}
}
