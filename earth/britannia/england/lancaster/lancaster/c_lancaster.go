package lancaster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LancasterCounty interface {
	feud.County
	BClitheroe克利瑟罗() feud.Barony
	BFurness弗内斯() feud.Barony
	BGisburn吉斯本() feud.Barony
	BLancaster兰开斯特() feud.Barony
	BPreston普雷斯顿() feud.Barony
	BSalford索尔福德() feud.Barony
	BSawley索利() feud.Barony
}

type 兰开斯特LancasterCounty struct {
	feud.BaseCounty
	克利瑟罗Clitheroe feud.Barony
	弗内斯Furness    feud.Barony
	吉斯本Gisburn    feud.Barony
	兰开斯特Lancaster feud.Barony
	普雷斯顿Preston   feud.Barony
	索尔福德Salford   feud.Barony
	索利Sawley      feud.Barony
}

func (c *兰开斯特LancasterCounty) BClitheroe克利瑟罗() feud.Barony {
	return c.克利瑟罗Clitheroe
}

func (c *兰开斯特LancasterCounty) BFurness弗内斯() feud.Barony {
	return c.弗内斯Furness
}

func (c *兰开斯特LancasterCounty) BGisburn吉斯本() feud.Barony {
	return c.吉斯本Gisburn
}

func (c *兰开斯特LancasterCounty) BLancaster兰开斯特() feud.Barony {
	return c.兰开斯特Lancaster
}

func (c *兰开斯特LancasterCounty) BPreston普雷斯顿() feud.Barony {
	return c.普雷斯顿Preston
}

func (c *兰开斯特LancasterCounty) BSalford索尔福德() feud.Barony {
	return c.索尔福德Salford
}

func (c *兰开斯特LancasterCounty) BSawley索利() feud.Barony {
	return c.索利Sawley
}

var CLancaster兰开斯特 LancasterCounty = &兰开斯特LancasterCounty{}

func init() {
	f := CLancaster兰开斯特.(*兰开斯特LancasterCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "58",
		Title:     "lancaster",
		TitleName: "兰开斯特",
		TitleCode: "c_lancaster",
		Baronies:  map[string]feud.Barony{},
	}

	f.克利瑟罗Clitheroe = BClitheroe克利瑟罗
	f.克利瑟罗Clitheroe.SetParent(f)

	f.弗内斯Furness = BFurness弗内斯
	f.弗内斯Furness.SetParent(f)

	f.吉斯本Gisburn = BGisburn吉斯本
	f.吉斯本Gisburn.SetParent(f)

	f.兰开斯特Lancaster = BLancaster兰开斯特
	f.兰开斯特Lancaster.SetParent(f)

	f.普雷斯顿Preston = BPreston普雷斯顿
	f.普雷斯顿Preston.SetParent(f)

	f.索尔福德Salford = BSalford索尔福德
	f.索尔福德Salford.SetParent(f)

	f.索利Sawley = BSawley索利
	f.索利Sawley.SetParent(f)

}
