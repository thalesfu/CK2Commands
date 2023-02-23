package ludrava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LudravaCounty interface {
	feud.County
	BJaisalmer贾沙梅尔() feud.Barony
	BLanela罗内罗() feud.Barony
	BLudrava律陀罗婆() feud.Barony
	BPhalavardhika颇罗伐地迦() feud.Barony
	BRanakpur罗那迦补罗() feud.Barony
	BRunjha蓝耆() feud.Barony
	BSahna犀诃那() feud.Barony
}

type 律陀罗婆LudravaCounty struct {
	feud.BaseCounty
	贾沙梅尔Jaisalmer      feud.Barony
	罗内罗Lanela          feud.Barony
	律陀罗婆Ludrava        feud.Barony
	颇罗伐地迦Phalavardhika feud.Barony
	罗那迦补罗Ranakpur      feud.Barony
	蓝耆Runjha           feud.Barony
	犀诃那Sahna           feud.Barony
}

func (c *律陀罗婆LudravaCounty) BJaisalmer贾沙梅尔() feud.Barony {
	return c.贾沙梅尔Jaisalmer
}

func (c *律陀罗婆LudravaCounty) BLanela罗内罗() feud.Barony {
	return c.罗内罗Lanela
}

func (c *律陀罗婆LudravaCounty) BLudrava律陀罗婆() feud.Barony {
	return c.律陀罗婆Ludrava
}

func (c *律陀罗婆LudravaCounty) BPhalavardhika颇罗伐地迦() feud.Barony {
	return c.颇罗伐地迦Phalavardhika
}

func (c *律陀罗婆LudravaCounty) BRanakpur罗那迦补罗() feud.Barony {
	return c.罗那迦补罗Ranakpur
}

func (c *律陀罗婆LudravaCounty) BRunjha蓝耆() feud.Barony {
	return c.蓝耆Runjha
}

func (c *律陀罗婆LudravaCounty) BSahna犀诃那() feud.Barony {
	return c.犀诃那Sahna
}

var CLudrava律陀罗婆 LudravaCounty = &律陀罗婆LudravaCounty{}

func init() {
	f := CLudrava律陀罗婆.(*律陀罗婆LudravaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1304",
		Title:     "ludrava",
		TitleName: "律陀罗婆",
		TitleCode: "c_ludrava",
		Baronies:  map[string]feud.Barony{},
	}

	f.贾沙梅尔Jaisalmer = BJaisalmer贾沙梅尔
	f.贾沙梅尔Jaisalmer.SetParent(f)

	f.罗内罗Lanela = BLanela罗内罗
	f.罗内罗Lanela.SetParent(f)

	f.律陀罗婆Ludrava = BLudrava律陀罗婆
	f.律陀罗婆Ludrava.SetParent(f)

	f.颇罗伐地迦Phalavardhika = BPhalavardhika颇罗伐地迦
	f.颇罗伐地迦Phalavardhika.SetParent(f)

	f.罗那迦补罗Ranakpur = BRanakpur罗那迦补罗
	f.罗那迦补罗Ranakpur.SetParent(f)

	f.蓝耆Runjha = BRunjha蓝耆
	f.蓝耆Runjha.SetParent(f)

	f.犀诃那Sahna = BSahna犀诃那
	f.犀诃那Sahna.SetParent(f)

}
