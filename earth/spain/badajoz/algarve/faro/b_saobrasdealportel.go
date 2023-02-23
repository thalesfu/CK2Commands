package faro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣布拉什迪阿尔波特尔SaobrasdealportelBarony struct {
	feud.BaseBarony
}

var BSaobrasdealportel圣布拉什迪阿尔波特尔 feud.Barony = &圣布拉什迪阿尔波特尔SaobrasdealportelBarony{}

func init() {
	f := BSaobrasdealportel圣布拉什迪阿尔波特尔.(*圣布拉什迪阿尔波特尔SaobrasdealportelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saobrasdealportel",
		TitleName: "圣布拉什迪阿尔波特尔",
		TitleCode: "b_saobrasdealportel",
	}
}
