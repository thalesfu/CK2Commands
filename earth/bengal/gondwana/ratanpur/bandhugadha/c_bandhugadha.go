package bandhugadha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BandhugadhaCounty interface {
	feud.County
	BBahoriband婆呼利般陀() feud.Barony
	BBandhugadha畔度伽荼() feud.Barony
	BMotibennur牟提旁奴() feud.Barony
	BMuhri牟诃梨() feud.Barony
	BNandwas难瓦斯() feud.Barony
	BShahdol沙赫多尔() feud.Barony
	BVirateshwar毗罗知湿伐罗() feud.Barony
}

type 畔度伽荼BandhugadhaCounty struct {
	feud.BaseCounty
	婆呼利般陀Bahoriband   feud.Barony
	畔度伽荼Bandhugadha   feud.Barony
	牟提旁奴Motibennur    feud.Barony
	牟诃梨Muhri          feud.Barony
	难瓦斯Nandwas        feud.Barony
	沙赫多尔Shahdol       feud.Barony
	毗罗知湿伐罗Virateshwar feud.Barony
}

func (c *畔度伽荼BandhugadhaCounty) BBahoriband婆呼利般陀() feud.Barony {
	return c.婆呼利般陀Bahoriband
}

func (c *畔度伽荼BandhugadhaCounty) BBandhugadha畔度伽荼() feud.Barony {
	return c.畔度伽荼Bandhugadha
}

func (c *畔度伽荼BandhugadhaCounty) BMotibennur牟提旁奴() feud.Barony {
	return c.牟提旁奴Motibennur
}

func (c *畔度伽荼BandhugadhaCounty) BMuhri牟诃梨() feud.Barony {
	return c.牟诃梨Muhri
}

func (c *畔度伽荼BandhugadhaCounty) BNandwas难瓦斯() feud.Barony {
	return c.难瓦斯Nandwas
}

func (c *畔度伽荼BandhugadhaCounty) BShahdol沙赫多尔() feud.Barony {
	return c.沙赫多尔Shahdol
}

func (c *畔度伽荼BandhugadhaCounty) BVirateshwar毗罗知湿伐罗() feud.Barony {
	return c.毗罗知湿伐罗Virateshwar
}

var CBandhugadha畔度伽荼 BandhugadhaCounty = &畔度伽荼BandhugadhaCounty{}

func init() {
	f := CBandhugadha畔度伽荼.(*畔度伽荼BandhugadhaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1165",
		Title:     "bandhugadha",
		TitleName: "畔度伽荼",
		TitleCode: "c_bandhugadha",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆呼利般陀Bahoriband = BBahoriband婆呼利般陀
	f.婆呼利般陀Bahoriband.SetParent(f)

	f.畔度伽荼Bandhugadha = BBandhugadha畔度伽荼
	f.畔度伽荼Bandhugadha.SetParent(f)

	f.牟提旁奴Motibennur = BMotibennur牟提旁奴
	f.牟提旁奴Motibennur.SetParent(f)

	f.牟诃梨Muhri = BMuhri牟诃梨
	f.牟诃梨Muhri.SetParent(f)

	f.难瓦斯Nandwas = BNandwas难瓦斯
	f.难瓦斯Nandwas.SetParent(f)

	f.沙赫多尔Shahdol = BShahdol沙赫多尔
	f.沙赫多尔Shahdol.SetParent(f)

	f.毗罗知湿伐罗Virateshwar = BVirateshwar毗罗知湿伐罗
	f.毗罗知湿伐罗Virateshwar.SetParent(f)

}
