package herakleia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗拉维奥波利斯FlaviopolisBarony struct {
	feud.BaseBarony
}

var BFlaviopolis弗拉维奥波利斯 feud.Barony = &弗拉维奥波利斯FlaviopolisBarony{}

func init() {
	f := BFlaviopolis弗拉维奥波利斯.(*弗拉维奥波利斯FlaviopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "flaviopolis",
		TitleName: "弗拉维奥波利斯",
		TitleCode: "b_flaviopolis",
	}
}
