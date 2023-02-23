package derby

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DerbyCounty interface {
	feud.County
	BBakewell贝克韦尔() feud.Barony
	BBolsover博尔索弗() feud.Barony
	BBurton伯顿() feud.Barony
	BCastleton卡斯尔顿() feud.Barony
	BChesterfield切斯特菲尔德() feud.Barony
	BDerby德比() feud.Barony
	BRepton雷普顿() feud.Barony
	BWirksworth威克斯沃斯() feud.Barony
}

type 德比DerbyCounty struct {
	feud.BaseCounty
	贝克韦尔Bakewell       feud.Barony
	博尔索弗Bolsover       feud.Barony
	伯顿Burton           feud.Barony
	卡斯尔顿Castleton      feud.Barony
	切斯特菲尔德Chesterfield feud.Barony
	德比Derby            feud.Barony
	雷普顿Repton          feud.Barony
	威克斯沃斯Wirksworth    feud.Barony
}

func (c *德比DerbyCounty) BBakewell贝克韦尔() feud.Barony {
	return c.贝克韦尔Bakewell
}

func (c *德比DerbyCounty) BBolsover博尔索弗() feud.Barony {
	return c.博尔索弗Bolsover
}

func (c *德比DerbyCounty) BBurton伯顿() feud.Barony {
	return c.伯顿Burton
}

func (c *德比DerbyCounty) BCastleton卡斯尔顿() feud.Barony {
	return c.卡斯尔顿Castleton
}

func (c *德比DerbyCounty) BChesterfield切斯特菲尔德() feud.Barony {
	return c.切斯特菲尔德Chesterfield
}

func (c *德比DerbyCounty) BDerby德比() feud.Barony {
	return c.德比Derby
}

func (c *德比DerbyCounty) BRepton雷普顿() feud.Barony {
	return c.雷普顿Repton
}

func (c *德比DerbyCounty) BWirksworth威克斯沃斯() feud.Barony {
	return c.威克斯沃斯Wirksworth
}

var CDerby德比 DerbyCounty = &德比DerbyCounty{}

func init() {
	f := CDerby德比.(*德比DerbyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "63",
		Title:     "derby",
		TitleName: "德比",
		TitleCode: "c_derby",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝克韦尔Bakewell = BBakewell贝克韦尔
	f.贝克韦尔Bakewell.SetParent(f)

	f.博尔索弗Bolsover = BBolsover博尔索弗
	f.博尔索弗Bolsover.SetParent(f)

	f.伯顿Burton = BBurton伯顿
	f.伯顿Burton.SetParent(f)

	f.卡斯尔顿Castleton = BCastleton卡斯尔顿
	f.卡斯尔顿Castleton.SetParent(f)

	f.切斯特菲尔德Chesterfield = BChesterfield切斯特菲尔德
	f.切斯特菲尔德Chesterfield.SetParent(f)

	f.德比Derby = BDerby德比
	f.德比Derby.SetParent(f)

	f.雷普顿Repton = BRepton雷普顿
	f.雷普顿Repton.SetParent(f)

	f.威克斯沃斯Wirksworth = BWirksworth威克斯沃斯
	f.威克斯沃斯Wirksworth.SetParent(f)

}
