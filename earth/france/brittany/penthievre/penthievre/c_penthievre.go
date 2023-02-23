package penthievre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PenthievreCounty interface {
	feud.County
	BChatelaudren沙泰洛德朗() feud.Barony
	BLoudeac卢代阿克() feud.Barony
	BMonkontour蒙孔图尔() feud.Barony
	BPaimpol潘波勒() feud.Barony
	BPeran佩朗() feud.Barony
	BQuintin坎坦() feud.Barony
	BVerdelet韦尔德莱() feud.Barony
}

type 庞蒂耶夫尔PenthievreCounty struct {
	feud.BaseCounty
	沙泰洛德朗Chatelaudren feud.Barony
	卢代阿克Loudeac       feud.Barony
	蒙孔图尔Monkontour    feud.Barony
	潘波勒Paimpol        feud.Barony
	佩朗Peran           feud.Barony
	坎坦Quintin         feud.Barony
	韦尔德莱Verdelet      feud.Barony
}

func (c *庞蒂耶夫尔PenthievreCounty) BChatelaudren沙泰洛德朗() feud.Barony {
	return c.沙泰洛德朗Chatelaudren
}

func (c *庞蒂耶夫尔PenthievreCounty) BLoudeac卢代阿克() feud.Barony {
	return c.卢代阿克Loudeac
}

func (c *庞蒂耶夫尔PenthievreCounty) BMonkontour蒙孔图尔() feud.Barony {
	return c.蒙孔图尔Monkontour
}

func (c *庞蒂耶夫尔PenthievreCounty) BPaimpol潘波勒() feud.Barony {
	return c.潘波勒Paimpol
}

func (c *庞蒂耶夫尔PenthievreCounty) BPeran佩朗() feud.Barony {
	return c.佩朗Peran
}

func (c *庞蒂耶夫尔PenthievreCounty) BQuintin坎坦() feud.Barony {
	return c.坎坦Quintin
}

func (c *庞蒂耶夫尔PenthievreCounty) BVerdelet韦尔德莱() feud.Barony {
	return c.韦尔德莱Verdelet
}

var CPenthievre庞蒂耶夫尔 PenthievreCounty = &庞蒂耶夫尔PenthievreCounty{}

func init() {
	f := CPenthievre庞蒂耶夫尔.(*庞蒂耶夫尔PenthievreCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "102",
		Title:     "penthievre",
		TitleName: "庞蒂耶夫尔",
		TitleCode: "c_penthievre",
		Baronies:  map[string]feud.Barony{},
	}

	f.沙泰洛德朗Chatelaudren = BChatelaudren沙泰洛德朗
	f.沙泰洛德朗Chatelaudren.SetParent(f)

	f.卢代阿克Loudeac = BLoudeac卢代阿克
	f.卢代阿克Loudeac.SetParent(f)

	f.蒙孔图尔Monkontour = BMonkontour蒙孔图尔
	f.蒙孔图尔Monkontour.SetParent(f)

	f.潘波勒Paimpol = BPaimpol潘波勒
	f.潘波勒Paimpol.SetParent(f)

	f.佩朗Peran = BPeran佩朗
	f.佩朗Peran.SetParent(f)

	f.坎坦Quintin = BQuintin坎坦
	f.坎坦Quintin.SetParent(f)

	f.韦尔德莱Verdelet = BVerdelet韦尔德莱
	f.韦尔德莱Verdelet.SetParent(f)

}
