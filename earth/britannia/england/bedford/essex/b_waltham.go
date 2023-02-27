package essex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔瑟姆WalthamBarony struct {
	feud.BaseBarony
}

var BWaltham沃尔瑟姆 feud.Barony = &沃尔瑟姆WalthamBarony{}

func init() {
    f := BWaltham沃尔瑟姆.(*沃尔瑟姆WalthamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "waltham",
		TitleName: "沃尔瑟姆",
		TitleCode: "b_waltham",
	}
}
