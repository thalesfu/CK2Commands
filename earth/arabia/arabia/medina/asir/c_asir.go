package asir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AsirCounty interface {
    feud.County
    BBahah巴哈() 	feud.Barony
    BDam达姆() 	feud.Barony
    BHajjah哈杰() 	feud.Barony
    BKamaran卡马兰() 	feud.Barony
    BKuthba呼图白() 	feud.Barony
    BSadah萨达() 	feud.Barony
    BTabala泰巴莱() 	feud.Barony
}

type 阿西尔AsirCounty struct {
	feud.BaseCounty
	巴哈Bahah 	feud.Barony
	达姆Dam 	feud.Barony
	哈杰Hajjah 	feud.Barony
	卡马兰Kamaran 	feud.Barony
	呼图白Kuthba 	feud.Barony
	萨达Sadah 	feud.Barony
	泰巴莱Tabala 	feud.Barony
}

func (c *阿西尔AsirCounty) BBahah巴哈() feud.Barony {
	return c.巴哈Bahah
}
    
func (c *阿西尔AsirCounty) BDam达姆() feud.Barony {
	return c.达姆Dam
}
    
func (c *阿西尔AsirCounty) BHajjah哈杰() feud.Barony {
	return c.哈杰Hajjah
}
    
func (c *阿西尔AsirCounty) BKamaran卡马兰() feud.Barony {
	return c.卡马兰Kamaran
}
    
func (c *阿西尔AsirCounty) BKuthba呼图白() feud.Barony {
	return c.呼图白Kuthba
}
    
func (c *阿西尔AsirCounty) BSadah萨达() feud.Barony {
	return c.萨达Sadah
}
    
func (c *阿西尔AsirCounty) BTabala泰巴莱() feud.Barony {
	return c.泰巴莱Tabala
}
    
var CAsir阿西尔 AsirCounty = &阿西尔AsirCounty{}

func init() {
	f := CAsir阿西尔.(*阿西尔AsirCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "861",
		Title:     "asir",
		TitleName: "阿西尔",
		TitleCode: "c_asir",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴哈Bahah = BBahah巴哈
	f.巴哈Bahah.SetParent(f)
	
	f.达姆Dam = BDam达姆
	f.达姆Dam.SetParent(f)
	
	f.哈杰Hajjah = BHajjah哈杰
	f.哈杰Hajjah.SetParent(f)
	
	f.卡马兰Kamaran = BKamaran卡马兰
	f.卡马兰Kamaran.SetParent(f)
	
	f.呼图白Kuthba = BKuthba呼图白
	f.呼图白Kuthba.SetParent(f)
	
	f.萨达Sadah = BSadah萨达
	f.萨达Sadah.SetParent(f)
	
	f.泰巴莱Tabala = BTabala泰巴莱
	f.泰巴莱Tabala.SetParent(f)
	
}
