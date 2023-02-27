package limbuwan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瞿揭罗诃GograhaBarony struct {
	feud.BaseBarony
}

var BGograha瞿揭罗诃 feud.Barony = &瞿揭罗诃GograhaBarony{}

func init() {
    f := BGograha瞿揭罗诃.(*瞿揭罗诃GograhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gograha",
		TitleName: "瞿揭罗诃",
		TitleCode: "b_gograha",
	}
}
