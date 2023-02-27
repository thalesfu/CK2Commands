package laodikeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗拉优波利斯FlaviupolisBarony struct {
	feud.BaseBarony
}

var BFlaviupolis弗拉优波利斯 feud.Barony = &弗拉优波利斯FlaviupolisBarony{}

func init() {
    f := BFlaviupolis弗拉优波利斯.(*弗拉优波利斯FlaviupolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "flaviupolis",
		TitleName: "弗拉优波利斯",
		TitleCode: "b_flaviupolis",
	}
}
