package slesvig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 石勒苏益格SlesvigBarony struct {
	feud.BaseBarony
}

var BSlesvig石勒苏益格 feud.Barony = &石勒苏益格SlesvigBarony{}

func init() {
    f := BSlesvig石勒苏益格.(*石勒苏益格SlesvigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "slesvig",
		TitleName: "石勒苏益格",
		TitleCode: "b_slesvig",
	}
}
