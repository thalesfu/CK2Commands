package trigarta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TrigartaCounty interface {
	feud.County
	BJalandhar阇烂达罗() feud.Barony
	BKaparthula迦补罗他罗() feud.Barony
	BMilwat密瓦特() feud.Barony
	BPathankot钵吒那拘吒() feud.Barony
	BRagnagar罗阇那揭罗() feud.Barony
	BRahon拉亨() feud.Barony
	BSarwmanpur萨婆摩那补罗() feud.Barony
}

type 三穴TrigartaCounty struct {
	feud.BaseCounty
	阇烂达罗Jalandhar    feud.Barony
	迦补罗他罗Kaparthula  feud.Barony
	密瓦特Milwat        feud.Barony
	钵吒那拘吒Pathankot   feud.Barony
	罗阇那揭罗Ragnagar    feud.Barony
	拉亨Rahon          feud.Barony
	萨婆摩那补罗Sarwmanpur feud.Barony
}

func (c *三穴TrigartaCounty) BJalandhar阇烂达罗() feud.Barony {
	return c.阇烂达罗Jalandhar
}

func (c *三穴TrigartaCounty) BKaparthula迦补罗他罗() feud.Barony {
	return c.迦补罗他罗Kaparthula
}

func (c *三穴TrigartaCounty) BMilwat密瓦特() feud.Barony {
	return c.密瓦特Milwat
}

func (c *三穴TrigartaCounty) BPathankot钵吒那拘吒() feud.Barony {
	return c.钵吒那拘吒Pathankot
}

func (c *三穴TrigartaCounty) BRagnagar罗阇那揭罗() feud.Barony {
	return c.罗阇那揭罗Ragnagar
}

func (c *三穴TrigartaCounty) BRahon拉亨() feud.Barony {
	return c.拉亨Rahon
}

func (c *三穴TrigartaCounty) BSarwmanpur萨婆摩那补罗() feud.Barony {
	return c.萨婆摩那补罗Sarwmanpur
}

var CTrigarta三穴 TrigartaCounty = &三穴TrigartaCounty{}

func init() {
	f := CTrigarta三穴.(*三穴TrigartaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1193",
		Title:     "trigarta",
		TitleName: "三穴",
		TitleCode: "c_trigarta",
		Baronies:  map[string]feud.Barony{},
	}

	f.阇烂达罗Jalandhar = BJalandhar阇烂达罗
	f.阇烂达罗Jalandhar.SetParent(f)

	f.迦补罗他罗Kaparthula = BKaparthula迦补罗他罗
	f.迦补罗他罗Kaparthula.SetParent(f)

	f.密瓦特Milwat = BMilwat密瓦特
	f.密瓦特Milwat.SetParent(f)

	f.钵吒那拘吒Pathankot = BPathankot钵吒那拘吒
	f.钵吒那拘吒Pathankot.SetParent(f)

	f.罗阇那揭罗Ragnagar = BRagnagar罗阇那揭罗
	f.罗阇那揭罗Ragnagar.SetParent(f)

	f.拉亨Rahon = BRahon拉亨
	f.拉亨Rahon.SetParent(f)

	f.萨婆摩那补罗Sarwmanpur = BSarwmanpur萨婆摩那补罗
	f.萨婆摩那补罗Sarwmanpur.SetParent(f)

}
