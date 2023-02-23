package baygal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蔑儿乞MerkitBarony struct {
	feud.BaseBarony
}

var BMerkit蔑儿乞 feud.Barony = &蔑儿乞MerkitBarony{}

func init() {
	f := BMerkit蔑儿乞.(*蔑儿乞MerkitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "merkit",
		TitleName: "蔑儿乞",
		TitleCode: "b_merkit",
	}
}
