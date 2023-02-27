package retz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉贝尔讷里La_bernerieBarony struct {
	feud.BaseBarony
}

var BLa_bernerie拉贝尔讷里 feud.Barony = &拉贝尔讷里La_bernerieBarony{}

func init() {
    f := BLa_bernerie拉贝尔讷里.(*拉贝尔讷里La_bernerieBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "la_bernerie",
		TitleName: "拉贝尔讷里",
		TitleCode: "b_la_bernerie",
	}
}
