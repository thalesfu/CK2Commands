package lakhnau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LakhnauCounty interface {
	feud.County
	BGaneshpur伽泥舍补罗() feud.Barony
	BGangabana恒伽婆那() feud.Barony
	BHardoi诃罗豆伊() feud.Barony
	BJasnaul耶娑那乌罗() feud.Barony
	BLakhnau落伽奈() feud.Barony
	BSatrikh娑帝力迦() feud.Barony
	BZaidpur阇耶陀城() feud.Barony
}

type 落伽奈LakhnauCounty struct {
	feud.BaseCounty
	伽泥舍补罗Ganeshpur feud.Barony
	恒伽婆那Gangabana  feud.Barony
	诃罗豆伊Hardoi     feud.Barony
	耶娑那乌罗Jasnaul   feud.Barony
	落伽奈Lakhnau     feud.Barony
	娑帝力迦Satrikh    feud.Barony
	阇耶陀城Zaidpur    feud.Barony
}

func (c *落伽奈LakhnauCounty) BGaneshpur伽泥舍补罗() feud.Barony {
	return c.伽泥舍补罗Ganeshpur
}

func (c *落伽奈LakhnauCounty) BGangabana恒伽婆那() feud.Barony {
	return c.恒伽婆那Gangabana
}

func (c *落伽奈LakhnauCounty) BHardoi诃罗豆伊() feud.Barony {
	return c.诃罗豆伊Hardoi
}

func (c *落伽奈LakhnauCounty) BJasnaul耶娑那乌罗() feud.Barony {
	return c.耶娑那乌罗Jasnaul
}

func (c *落伽奈LakhnauCounty) BLakhnau落伽奈() feud.Barony {
	return c.落伽奈Lakhnau
}

func (c *落伽奈LakhnauCounty) BSatrikh娑帝力迦() feud.Barony {
	return c.娑帝力迦Satrikh
}

func (c *落伽奈LakhnauCounty) BZaidpur阇耶陀城() feud.Barony {
	return c.阇耶陀城Zaidpur
}

var CLakhnau落伽奈 LakhnauCounty = &落伽奈LakhnauCounty{}

func init() {
	f := CLakhnau落伽奈.(*落伽奈LakhnauCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1167",
		Title:     "lakhnau",
		TitleName: "落伽奈",
		TitleCode: "c_lakhnau",
		Baronies:  map[string]feud.Barony{},
	}

	f.伽泥舍补罗Ganeshpur = BGaneshpur伽泥舍补罗
	f.伽泥舍补罗Ganeshpur.SetParent(f)

	f.恒伽婆那Gangabana = BGangabana恒伽婆那
	f.恒伽婆那Gangabana.SetParent(f)

	f.诃罗豆伊Hardoi = BHardoi诃罗豆伊
	f.诃罗豆伊Hardoi.SetParent(f)

	f.耶娑那乌罗Jasnaul = BJasnaul耶娑那乌罗
	f.耶娑那乌罗Jasnaul.SetParent(f)

	f.落伽奈Lakhnau = BLakhnau落伽奈
	f.落伽奈Lakhnau.SetParent(f)

	f.娑帝力迦Satrikh = BSatrikh娑帝力迦
	f.娑帝力迦Satrikh.SetParent(f)

	f.阇耶陀城Zaidpur = BZaidpur阇耶陀城
	f.阇耶陀城Zaidpur.SetParent(f)

}
