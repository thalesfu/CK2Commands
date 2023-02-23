package ancona

import (
	"github.com/thalesfu/CK2Commands/earth/italy/papacy/ancona/ancona"
	"github.com/thalesfu/CK2Commands/earth/italy/papacy/ancona/rimini"
	"github.com/thalesfu/CK2Commands/earth/italy/papacy/ancona/urbino"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AnconaDuke interface {
	feud.Duke
	CAncona安科纳() ancona.AnconaCounty
	CRimini里米尼() rimini.RiminiCounty
	CUrbino乌尔比诺() urbino.UrbinoCounty
}

type 安科纳AnconaDuke struct {
	feud.BaseDuke
	安科纳Ancona  ancona.AnconaCounty
	里米尼Rimini  rimini.RiminiCounty
	乌尔比诺Urbino urbino.UrbinoCounty
}

func (d *安科纳AnconaDuke) CAncona安科纳() ancona.AnconaCounty {
	return d.安科纳Ancona
}

func (d *安科纳AnconaDuke) CRimini里米尼() rimini.RiminiCounty {
	return d.里米尼Rimini
}

func (d *安科纳AnconaDuke) CUrbino乌尔比诺() urbino.UrbinoCounty {
	return d.乌尔比诺Urbino
}

var DAncona安科纳 AnconaDuke = &安科纳AnconaDuke{}

func init() {
	f := DAncona安科纳.(*安科纳AnconaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ancona",
		TitleName: "安科纳",
		TitleCode: "d_ancona",
		Counties:  map[string]feud.County{},
	}

	f.安科纳Ancona = ancona.CAncona安科纳
	f.安科纳Ancona.SetParent(f)

	f.里米尼Rimini = rimini.CRimini里米尼
	f.里米尼Rimini.SetParent(f)

	f.乌尔比诺Urbino = urbino.CUrbino乌尔比诺
	f.乌尔比诺Urbino.SetParent(f)

}
