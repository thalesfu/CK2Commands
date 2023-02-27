package maragha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼苏尔卢MansurluBarony struct {
	feud.BaseBarony
}

var BMansurlu曼苏尔卢 feud.Barony = &曼苏尔卢MansurluBarony{}

func init() {
    f := BMansurlu曼苏尔卢.(*曼苏尔卢MansurluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mansurlu",
		TitleName: "曼苏尔卢",
		TitleCode: "b_mansurlu",
	}
}
