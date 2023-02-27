package aland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 松德SundBarony struct {
	feud.BaseBarony
}

var BSund松德 feud.Barony = &松德SundBarony{}

func init() {
    f := BSund松德.(*松德SundBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sund",
		TitleName: "松德",
		TitleCode: "b_sund",
	}
}
