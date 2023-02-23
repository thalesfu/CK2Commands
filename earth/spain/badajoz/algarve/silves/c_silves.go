package silves

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SilvesCounty interface {
	feud.County
	BAljezur阿尔热祖尔() feud.Barony
	BAlmodovar阿尔莫多瓦() feud.Barony
	BAlvor阿沃尔() feud.Barony
	BArrifana阿里法纳() feud.Barony
	BLagos拉古什() feud.Barony
	BMonchique蒙希克() feud.Barony
	BSilves锡尔维什() feud.Barony
}

type 锡尔维什SilvesCounty struct {
	feud.BaseCounty
	阿尔热祖尔Aljezur   feud.Barony
	阿尔莫多瓦Almodovar feud.Barony
	阿沃尔Alvor       feud.Barony
	阿里法纳Arrifana   feud.Barony
	拉古什Lagos       feud.Barony
	蒙希克Monchique   feud.Barony
	锡尔维什Silves     feud.Barony
}

func (c *锡尔维什SilvesCounty) BAljezur阿尔热祖尔() feud.Barony {
	return c.阿尔热祖尔Aljezur
}

func (c *锡尔维什SilvesCounty) BAlmodovar阿尔莫多瓦() feud.Barony {
	return c.阿尔莫多瓦Almodovar
}

func (c *锡尔维什SilvesCounty) BAlvor阿沃尔() feud.Barony {
	return c.阿沃尔Alvor
}

func (c *锡尔维什SilvesCounty) BArrifana阿里法纳() feud.Barony {
	return c.阿里法纳Arrifana
}

func (c *锡尔维什SilvesCounty) BLagos拉古什() feud.Barony {
	return c.拉古什Lagos
}

func (c *锡尔维什SilvesCounty) BMonchique蒙希克() feud.Barony {
	return c.蒙希克Monchique
}

func (c *锡尔维什SilvesCounty) BSilves锡尔维什() feud.Barony {
	return c.锡尔维什Silves
}

var CSilves锡尔维什 SilvesCounty = &锡尔维什SilvesCounty{}

func init() {
	f := CSilves锡尔维什.(*锡尔维什SilvesCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "162",
		Title:     "silves",
		TitleName: "锡尔维什",
		TitleCode: "c_silves",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔热祖尔Aljezur = BAljezur阿尔热祖尔
	f.阿尔热祖尔Aljezur.SetParent(f)

	f.阿尔莫多瓦Almodovar = BAlmodovar阿尔莫多瓦
	f.阿尔莫多瓦Almodovar.SetParent(f)

	f.阿沃尔Alvor = BAlvor阿沃尔
	f.阿沃尔Alvor.SetParent(f)

	f.阿里法纳Arrifana = BArrifana阿里法纳
	f.阿里法纳Arrifana.SetParent(f)

	f.拉古什Lagos = BLagos拉古什
	f.拉古什Lagos.SetParent(f)

	f.蒙希克Monchique = BMonchique蒙希克
	f.蒙希克Monchique.SetParent(f)

	f.锡尔维什Silves = BSilves锡尔维什
	f.锡尔维什Silves.SetParent(f)

}
