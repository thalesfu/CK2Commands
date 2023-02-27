package vexin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔库尔HarcourtBarony struct {
	feud.BaseBarony
}

var BHarcourt阿尔库尔 feud.Barony = &阿尔库尔HarcourtBarony{}

func init() {
    f := BHarcourt阿尔库尔.(*阿尔库尔HarcourtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harcourt",
		TitleName: "阿尔库尔",
		TitleCode: "b_harcourt",
	}
}
