package dasapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DasapuraCounty interface {
    feud.County
    BChinchwad真遮婆荼() 	feud.Barony
    BDalu陀卢() 	feud.Barony
    BDasapura陀舍补罗() 	feud.Barony
    BDharmrajeshwar达摩罗逝湿伐罗() 	feud.Barony
    BHinglajgarh兴具罗阇姞利呬() 	feud.Barony
    BSondani孙陀尼() 	feud.Barony
    BTaxakeshwar德叉计湿伐罗() 	feud.Barony
}

type 陀舍补罗DasapuraCounty struct {
	feud.BaseCounty
	真遮婆荼Chinchwad 	feud.Barony
	陀卢Dalu 	feud.Barony
	陀舍补罗Dasapura 	feud.Barony
	达摩罗逝湿伐罗Dharmrajeshwar 	feud.Barony
	兴具罗阇姞利呬Hinglajgarh 	feud.Barony
	孙陀尼Sondani 	feud.Barony
	德叉计湿伐罗Taxakeshwar 	feud.Barony
}

func (c *陀舍补罗DasapuraCounty) BChinchwad真遮婆荼() feud.Barony {
	return c.真遮婆荼Chinchwad
}
    
func (c *陀舍补罗DasapuraCounty) BDalu陀卢() feud.Barony {
	return c.陀卢Dalu
}
    
func (c *陀舍补罗DasapuraCounty) BDasapura陀舍补罗() feud.Barony {
	return c.陀舍补罗Dasapura
}
    
func (c *陀舍补罗DasapuraCounty) BDharmrajeshwar达摩罗逝湿伐罗() feud.Barony {
	return c.达摩罗逝湿伐罗Dharmrajeshwar
}
    
func (c *陀舍补罗DasapuraCounty) BHinglajgarh兴具罗阇姞利呬() feud.Barony {
	return c.兴具罗阇姞利呬Hinglajgarh
}
    
func (c *陀舍补罗DasapuraCounty) BSondani孙陀尼() feud.Barony {
	return c.孙陀尼Sondani
}
    
func (c *陀舍补罗DasapuraCounty) BTaxakeshwar德叉计湿伐罗() feud.Barony {
	return c.德叉计湿伐罗Taxakeshwar
}
    
var CDasapura陀舍补罗 DasapuraCounty = &陀舍补罗DasapuraCounty{}

func init() {
	f := CDasapura陀舍补罗.(*陀舍补罗DasapuraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1148",
		Title:     "dasapura",
		TitleName: "陀舍补罗",
		TitleCode: "c_dasapura",
		Baronies:  map[string]feud.Barony{},
	}

	f.真遮婆荼Chinchwad = BChinchwad真遮婆荼
	f.真遮婆荼Chinchwad.SetParent(f)
	
	f.陀卢Dalu = BDalu陀卢
	f.陀卢Dalu.SetParent(f)
	
	f.陀舍补罗Dasapura = BDasapura陀舍补罗
	f.陀舍补罗Dasapura.SetParent(f)
	
	f.达摩罗逝湿伐罗Dharmrajeshwar = BDharmrajeshwar达摩罗逝湿伐罗
	f.达摩罗逝湿伐罗Dharmrajeshwar.SetParent(f)
	
	f.兴具罗阇姞利呬Hinglajgarh = BHinglajgarh兴具罗阇姞利呬
	f.兴具罗阇姞利呬Hinglajgarh.SetParent(f)
	
	f.孙陀尼Sondani = BSondani孙陀尼
	f.孙陀尼Sondani.SetParent(f)
	
	f.德叉计湿伐罗Taxakeshwar = BTaxakeshwar德叉计湿伐罗
	f.德叉计湿伐罗Taxakeshwar.SetParent(f)
	
}
