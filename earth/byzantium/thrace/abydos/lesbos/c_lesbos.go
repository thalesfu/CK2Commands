package lesbos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LesbosCounty interface {
    feud.County
    BAgiasos阿亚索斯() 	feud.Barony
    BEresos埃雷索斯() 	feud.Barony
    BKalloni卡洛尼() 	feud.Barony
    BMithymna米西姆纳() 	feud.Barony
    BMoudros穆兹罗斯() 	feud.Barony
    BMytilene米蒂利尼() 	feud.Barony
    BPlomari普洛马里() 	feud.Barony
    BThasos萨索斯() 	feud.Barony
}

type 莱斯沃斯LesbosCounty struct {
	feud.BaseCounty
	阿亚索斯Agiasos 	feud.Barony
	埃雷索斯Eresos 	feud.Barony
	卡洛尼Kalloni 	feud.Barony
	米西姆纳Mithymna 	feud.Barony
	穆兹罗斯Moudros 	feud.Barony
	米蒂利尼Mytilene 	feud.Barony
	普洛马里Plomari 	feud.Barony
	萨索斯Thasos 	feud.Barony
}

func (c *莱斯沃斯LesbosCounty) BAgiasos阿亚索斯() feud.Barony {
	return c.阿亚索斯Agiasos
}
    
func (c *莱斯沃斯LesbosCounty) BEresos埃雷索斯() feud.Barony {
	return c.埃雷索斯Eresos
}
    
func (c *莱斯沃斯LesbosCounty) BKalloni卡洛尼() feud.Barony {
	return c.卡洛尼Kalloni
}
    
func (c *莱斯沃斯LesbosCounty) BMithymna米西姆纳() feud.Barony {
	return c.米西姆纳Mithymna
}
    
func (c *莱斯沃斯LesbosCounty) BMoudros穆兹罗斯() feud.Barony {
	return c.穆兹罗斯Moudros
}
    
func (c *莱斯沃斯LesbosCounty) BMytilene米蒂利尼() feud.Barony {
	return c.米蒂利尼Mytilene
}
    
func (c *莱斯沃斯LesbosCounty) BPlomari普洛马里() feud.Barony {
	return c.普洛马里Plomari
}
    
func (c *莱斯沃斯LesbosCounty) BThasos萨索斯() feud.Barony {
	return c.萨索斯Thasos
}
    
var CLesbos莱斯沃斯 LesbosCounty = &莱斯沃斯LesbosCounty{}

func init() {
	f := CLesbos莱斯沃斯.(*莱斯沃斯LesbosCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "487",
		Title:     "lesbos",
		TitleName: "莱斯沃斯",
		TitleCode: "c_lesbos",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿亚索斯Agiasos = BAgiasos阿亚索斯
	f.阿亚索斯Agiasos.SetParent(f)
	
	f.埃雷索斯Eresos = BEresos埃雷索斯
	f.埃雷索斯Eresos.SetParent(f)
	
	f.卡洛尼Kalloni = BKalloni卡洛尼
	f.卡洛尼Kalloni.SetParent(f)
	
	f.米西姆纳Mithymna = BMithymna米西姆纳
	f.米西姆纳Mithymna.SetParent(f)
	
	f.穆兹罗斯Moudros = BMoudros穆兹罗斯
	f.穆兹罗斯Moudros.SetParent(f)
	
	f.米蒂利尼Mytilene = BMytilene米蒂利尼
	f.米蒂利尼Mytilene.SetParent(f)
	
	f.普洛马里Plomari = BPlomari普洛马里
	f.普洛马里Plomari.SetParent(f)
	
	f.萨索斯Thasos = BThasos萨索斯
	f.萨索斯Thasos.SetParent(f)
	
}
