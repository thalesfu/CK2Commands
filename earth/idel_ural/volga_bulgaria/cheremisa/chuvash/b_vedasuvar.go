package chuvash

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维达苏瓦尔VedasuvarBarony struct {
	feud.BaseBarony
}

var BVedasuvar维达苏瓦尔 feud.Barony = &维达苏瓦尔VedasuvarBarony{}

func init() {
    f := BVedasuvar维达苏瓦尔.(*维达苏瓦尔VedasuvarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vedasuvar",
		TitleName: "维达苏瓦尔",
		TitleCode: "b_vedasuvar",
	}
}
