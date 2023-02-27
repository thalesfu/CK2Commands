package vairata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 均均乌恩JhunjhunuBarony struct {
	feud.BaseBarony
}

var BJhunjhunu均均乌恩 feud.Barony = &均均乌恩JhunjhunuBarony{}

func init() {
    f := BJhunjhunu均均乌恩.(*均均乌恩JhunjhunuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jhunjhunu",
		TitleName: "均均乌恩",
		TitleCode: "b_jhunjhunu",
	}
}
