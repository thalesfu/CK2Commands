package algeciras

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlgecirasCounty interface {
	feud.County
	BAlgericas阿尔赫西拉斯() feud.Barony
	BCasares卡萨雷斯() feud.Barony
	BEstepona埃斯特波纳() feud.Barony
	BGibraltar直布罗陀() feud.Barony
	BJimenadelafrontera希梅纳德拉弗龙特拉() feud.Barony
	BRonda龙达() feud.Barony
	BSanroque圣罗克() feud.Barony
	BTarifa塔里法() feud.Barony
}

type 阿尔赫西拉斯AlgecirasCounty struct {
	feud.BaseCounty
	阿尔赫西拉斯Algericas             feud.Barony
	卡萨雷斯Casares                 feud.Barony
	埃斯特波纳Estepona               feud.Barony
	直布罗陀Gibraltar               feud.Barony
	希梅纳德拉弗龙特拉Jimenadelafrontera feud.Barony
	龙达Ronda                     feud.Barony
	圣罗克Sanroque                 feud.Barony
	塔里法Tarifa                   feud.Barony
}

func (c *阿尔赫西拉斯AlgecirasCounty) BAlgericas阿尔赫西拉斯() feud.Barony {
	return c.阿尔赫西拉斯Algericas
}

func (c *阿尔赫西拉斯AlgecirasCounty) BCasares卡萨雷斯() feud.Barony {
	return c.卡萨雷斯Casares
}

func (c *阿尔赫西拉斯AlgecirasCounty) BEstepona埃斯特波纳() feud.Barony {
	return c.埃斯特波纳Estepona
}

func (c *阿尔赫西拉斯AlgecirasCounty) BGibraltar直布罗陀() feud.Barony {
	return c.直布罗陀Gibraltar
}

func (c *阿尔赫西拉斯AlgecirasCounty) BJimenadelafrontera希梅纳德拉弗龙特拉() feud.Barony {
	return c.希梅纳德拉弗龙特拉Jimenadelafrontera
}

func (c *阿尔赫西拉斯AlgecirasCounty) BRonda龙达() feud.Barony {
	return c.龙达Ronda
}

func (c *阿尔赫西拉斯AlgecirasCounty) BSanroque圣罗克() feud.Barony {
	return c.圣罗克Sanroque
}

func (c *阿尔赫西拉斯AlgecirasCounty) BTarifa塔里法() feud.Barony {
	return c.塔里法Tarifa
}

var CAlgeciras阿尔赫西拉斯 AlgecirasCounty = &阿尔赫西拉斯AlgecirasCounty{}

func init() {
	f := CAlgeciras阿尔赫西拉斯.(*阿尔赫西拉斯AlgecirasCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "166",
		Title:     "algeciras",
		TitleName: "阿尔赫西拉斯",
		TitleCode: "c_algeciras",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔赫西拉斯Algericas = BAlgericas阿尔赫西拉斯
	f.阿尔赫西拉斯Algericas.SetParent(f)

	f.卡萨雷斯Casares = BCasares卡萨雷斯
	f.卡萨雷斯Casares.SetParent(f)

	f.埃斯特波纳Estepona = BEstepona埃斯特波纳
	f.埃斯特波纳Estepona.SetParent(f)

	f.直布罗陀Gibraltar = BGibraltar直布罗陀
	f.直布罗陀Gibraltar.SetParent(f)

	f.希梅纳德拉弗龙特拉Jimenadelafrontera = BJimenadelafrontera希梅纳德拉弗龙特拉
	f.希梅纳德拉弗龙特拉Jimenadelafrontera.SetParent(f)

	f.龙达Ronda = BRonda龙达
	f.龙达Ronda.SetParent(f)

	f.圣罗克Sanroque = BSanroque圣罗克
	f.圣罗克Sanroque.SetParent(f)

	f.塔里法Tarifa = BTarifa塔里法
	f.塔里法Tarifa.SetParent(f)

}
