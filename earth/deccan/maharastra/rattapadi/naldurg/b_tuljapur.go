package naldurg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图尔贾普尔TuljapurBarony struct {
	feud.BaseBarony
}

var BTuljapur图尔贾普尔 feud.Barony = &图尔贾普尔TuljapurBarony{}

func init() {
    f := BTuljapur图尔贾普尔.(*图尔贾普尔TuljapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tuljapur",
		TitleName: "图尔贾普尔",
		TitleCode: "b_tuljapur",
	}
}
