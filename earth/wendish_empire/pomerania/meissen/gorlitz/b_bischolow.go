package gorlitz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比绍洛夫BischolowBarony struct {
	feud.BaseBarony
}

var BBischolow比绍洛夫 feud.Barony = &比绍洛夫BischolowBarony{}

func init() {
    f := BBischolow比绍洛夫.(*比绍洛夫BischolowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bischolow",
		TitleName: "比绍洛夫",
		TitleCode: "b_bischolow",
	}
}
