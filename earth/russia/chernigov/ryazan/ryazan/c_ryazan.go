package ryazan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RyazanCounty interface {
    feud.County
    BKrasnoye克拉斯诺耶() 	feud.Barony
    BLukhovitsy卢霍维齐() 	feud.Barony
    BNovoselki诺沃肖尔基() 	feud.Barony
    BPereyaslavl_ryazanski佩列亚斯拉夫尔梁赞斯基() 	feud.Barony
    BRyazan梁赞() 	feud.Barony
    BRybnoye雷布诺耶() 	feud.Barony
    BTurlatovo图尔拉托沃() 	feud.Barony
}

type 梁赞RyazanCounty struct {
	feud.BaseCounty
	克拉斯诺耶Krasnoye 	feud.Barony
	卢霍维齐Lukhovitsy 	feud.Barony
	诺沃肖尔基Novoselki 	feud.Barony
	佩列亚斯拉夫尔梁赞斯基Pereyaslavl_ryazanski 	feud.Barony
	梁赞Ryazan 	feud.Barony
	雷布诺耶Rybnoye 	feud.Barony
	图尔拉托沃Turlatovo 	feud.Barony
}

func (c *梁赞RyazanCounty) BKrasnoye克拉斯诺耶() feud.Barony {
	return c.克拉斯诺耶Krasnoye
}
    
func (c *梁赞RyazanCounty) BLukhovitsy卢霍维齐() feud.Barony {
	return c.卢霍维齐Lukhovitsy
}
    
func (c *梁赞RyazanCounty) BNovoselki诺沃肖尔基() feud.Barony {
	return c.诺沃肖尔基Novoselki
}
    
func (c *梁赞RyazanCounty) BPereyaslavl_ryazanski佩列亚斯拉夫尔梁赞斯基() feud.Barony {
	return c.佩列亚斯拉夫尔梁赞斯基Pereyaslavl_ryazanski
}
    
func (c *梁赞RyazanCounty) BRyazan梁赞() feud.Barony {
	return c.梁赞Ryazan
}
    
func (c *梁赞RyazanCounty) BRybnoye雷布诺耶() feud.Barony {
	return c.雷布诺耶Rybnoye
}
    
func (c *梁赞RyazanCounty) BTurlatovo图尔拉托沃() feud.Barony {
	return c.图尔拉托沃Turlatovo
}
    
var CRyazan梁赞 RyazanCounty = &梁赞RyazanCounty{}

func init() {
	f := CRyazan梁赞.(*梁赞RyazanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "580",
		Title:     "ryazan",
		TitleName: "梁赞",
		TitleCode: "c_ryazan",
		Baronies:  map[string]feud.Barony{},
	}

	f.克拉斯诺耶Krasnoye = BKrasnoye克拉斯诺耶
	f.克拉斯诺耶Krasnoye.SetParent(f)
	
	f.卢霍维齐Lukhovitsy = BLukhovitsy卢霍维齐
	f.卢霍维齐Lukhovitsy.SetParent(f)
	
	f.诺沃肖尔基Novoselki = BNovoselki诺沃肖尔基
	f.诺沃肖尔基Novoselki.SetParent(f)
	
	f.佩列亚斯拉夫尔梁赞斯基Pereyaslavl_ryazanski = BPereyaslavl_ryazanski佩列亚斯拉夫尔梁赞斯基
	f.佩列亚斯拉夫尔梁赞斯基Pereyaslavl_ryazanski.SetParent(f)
	
	f.梁赞Ryazan = BRyazan梁赞
	f.梁赞Ryazan.SetParent(f)
	
	f.雷布诺耶Rybnoye = BRybnoye雷布诺耶
	f.雷布诺耶Rybnoye.SetParent(f)
	
	f.图尔拉托沃Turlatovo = BTurlatovo图尔拉托沃
	f.图尔拉托沃Turlatovo.SetParent(f)
	
}
