package suzdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SuzdalCounty interface {
    feud.County
    BIvanovo伊万诺沃() 	feud.Barony
    BKineshma基涅什马() 	feud.Barony
    BSereda谢列达() 	feud.Barony
    BShuya舒亚() 	feud.Barony
    BSudalyvka苏达雷夫卡() 	feud.Barony
    BSuzdal苏兹达尔() 	feud.Barony
    BVichuga维丘加() 	feud.Barony
}

type 苏兹达尔SuzdalCounty struct {
	feud.BaseCounty
	伊万诺沃Ivanovo 	feud.Barony
	基涅什马Kineshma 	feud.Barony
	谢列达Sereda 	feud.Barony
	舒亚Shuya 	feud.Barony
	苏达雷夫卡Sudalyvka 	feud.Barony
	苏兹达尔Suzdal 	feud.Barony
	维丘加Vichuga 	feud.Barony
}

func (c *苏兹达尔SuzdalCounty) BIvanovo伊万诺沃() feud.Barony {
	return c.伊万诺沃Ivanovo
}
    
func (c *苏兹达尔SuzdalCounty) BKineshma基涅什马() feud.Barony {
	return c.基涅什马Kineshma
}
    
func (c *苏兹达尔SuzdalCounty) BSereda谢列达() feud.Barony {
	return c.谢列达Sereda
}
    
func (c *苏兹达尔SuzdalCounty) BShuya舒亚() feud.Barony {
	return c.舒亚Shuya
}
    
func (c *苏兹达尔SuzdalCounty) BSudalyvka苏达雷夫卡() feud.Barony {
	return c.苏达雷夫卡Sudalyvka
}
    
func (c *苏兹达尔SuzdalCounty) BSuzdal苏兹达尔() feud.Barony {
	return c.苏兹达尔Suzdal
}
    
func (c *苏兹达尔SuzdalCounty) BVichuga维丘加() feud.Barony {
	return c.维丘加Vichuga
}
    
var CSuzdal苏兹达尔 SuzdalCounty = &苏兹达尔SuzdalCounty{}

func init() {
	f := CSuzdal苏兹达尔.(*苏兹达尔SuzdalCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "583",
		Title:     "suzdal",
		TitleName: "苏兹达尔",
		TitleCode: "c_suzdal",
		Baronies:  map[string]feud.Barony{},
	}

	f.伊万诺沃Ivanovo = BIvanovo伊万诺沃
	f.伊万诺沃Ivanovo.SetParent(f)
	
	f.基涅什马Kineshma = BKineshma基涅什马
	f.基涅什马Kineshma.SetParent(f)
	
	f.谢列达Sereda = BSereda谢列达
	f.谢列达Sereda.SetParent(f)
	
	f.舒亚Shuya = BShuya舒亚
	f.舒亚Shuya.SetParent(f)
	
	f.苏达雷夫卡Sudalyvka = BSudalyvka苏达雷夫卡
	f.苏达雷夫卡Sudalyvka.SetParent(f)
	
	f.苏兹达尔Suzdal = BSuzdal苏兹达尔
	f.苏兹达尔Suzdal.SetParent(f)
	
	f.维丘加Vichuga = BVichuga维丘加
	f.维丘加Vichuga.SetParent(f)
	
}
