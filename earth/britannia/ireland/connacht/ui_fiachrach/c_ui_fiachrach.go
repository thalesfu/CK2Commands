package ui_fiachrach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Ui_fiachrachCounty interface {
    feud.County
    BAughagower阿哈乔韦尔() 	feud.Barony
    BBallintubber巴林托伯() 	feud.Barony
    BCong康镇() 	feud.Barony
    BErrew埃拉() 	feud.Barony
    BInishkea伊尼什盖() 	feud.Barony
    BMayo梅奥() 	feud.Barony
    BMullet穆雷特() 	feud.Barony
}

type 北康诺特Ui_fiachrachCounty struct {
	feud.BaseCounty
	阿哈乔韦尔Aughagower 	feud.Barony
	巴林托伯Ballintubber 	feud.Barony
	康镇Cong 	feud.Barony
	埃拉Errew 	feud.Barony
	伊尼什盖Inishkea 	feud.Barony
	梅奥Mayo 	feud.Barony
	穆雷特Mullet 	feud.Barony
}

func (c *北康诺特Ui_fiachrachCounty) BAughagower阿哈乔韦尔() feud.Barony {
	return c.阿哈乔韦尔Aughagower
}
    
func (c *北康诺特Ui_fiachrachCounty) BBallintubber巴林托伯() feud.Barony {
	return c.巴林托伯Ballintubber
}
    
func (c *北康诺特Ui_fiachrachCounty) BCong康镇() feud.Barony {
	return c.康镇Cong
}
    
func (c *北康诺特Ui_fiachrachCounty) BErrew埃拉() feud.Barony {
	return c.埃拉Errew
}
    
func (c *北康诺特Ui_fiachrachCounty) BInishkea伊尼什盖() feud.Barony {
	return c.伊尼什盖Inishkea
}
    
func (c *北康诺特Ui_fiachrachCounty) BMayo梅奥() feud.Barony {
	return c.梅奥Mayo
}
    
func (c *北康诺特Ui_fiachrachCounty) BMullet穆雷特() feud.Barony {
	return c.穆雷特Mullet
}
    
var CUi_fiachrach北康诺特 Ui_fiachrachCounty = &北康诺特Ui_fiachrachCounty{}

func init() {
	f := CUi_fiachrach北康诺特.(*北康诺特Ui_fiachrachCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1954",
		Title:     "ui_fiachrach",
		TitleName: "北康诺特",
		TitleCode: "c_ui_fiachrach",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿哈乔韦尔Aughagower = BAughagower阿哈乔韦尔
	f.阿哈乔韦尔Aughagower.SetParent(f)
	
	f.巴林托伯Ballintubber = BBallintubber巴林托伯
	f.巴林托伯Ballintubber.SetParent(f)
	
	f.康镇Cong = BCong康镇
	f.康镇Cong.SetParent(f)
	
	f.埃拉Errew = BErrew埃拉
	f.埃拉Errew.SetParent(f)
	
	f.伊尼什盖Inishkea = BInishkea伊尼什盖
	f.伊尼什盖Inishkea.SetParent(f)
	
	f.梅奥Mayo = BMayo梅奥
	f.梅奥Mayo.SetParent(f)
	
	f.穆雷特Mullet = BMullet穆雷特
	f.穆雷特Mullet.SetParent(f)
	
}
