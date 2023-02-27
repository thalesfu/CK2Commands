package danzig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DanzigCounty interface {
    feud.County
    BDanlauenburg劳恩堡() 	feud.Barony
    BDanzig但泽() 	feud.Barony
    BMewe梅韦() 	feud.Barony
    BOliva奥利瓦() 	feud.Barony
    BPuck普茨克() 	feud.Barony
    BSchlochau施洛豪() 	feud.Barony
    BSchwetz施韦茨() 	feud.Barony
    BTuchel图赫尔() 	feud.Barony
}

type 但泽DanzigCounty struct {
	feud.BaseCounty
	劳恩堡Danlauenburg 	feud.Barony
	但泽Danzig 	feud.Barony
	梅韦Mewe 	feud.Barony
	奥利瓦Oliva 	feud.Barony
	普茨克Puck 	feud.Barony
	施洛豪Schlochau 	feud.Barony
	施韦茨Schwetz 	feud.Barony
	图赫尔Tuchel 	feud.Barony
}

func (c *但泽DanzigCounty) BDanlauenburg劳恩堡() feud.Barony {
	return c.劳恩堡Danlauenburg
}
    
func (c *但泽DanzigCounty) BDanzig但泽() feud.Barony {
	return c.但泽Danzig
}
    
func (c *但泽DanzigCounty) BMewe梅韦() feud.Barony {
	return c.梅韦Mewe
}
    
func (c *但泽DanzigCounty) BOliva奥利瓦() feud.Barony {
	return c.奥利瓦Oliva
}
    
func (c *但泽DanzigCounty) BPuck普茨克() feud.Barony {
	return c.普茨克Puck
}
    
func (c *但泽DanzigCounty) BSchlochau施洛豪() feud.Barony {
	return c.施洛豪Schlochau
}
    
func (c *但泽DanzigCounty) BSchwetz施韦茨() feud.Barony {
	return c.施韦茨Schwetz
}
    
func (c *但泽DanzigCounty) BTuchel图赫尔() feud.Barony {
	return c.图赫尔Tuchel
}
    
var CDanzig但泽 DanzigCounty = &但泽DanzigCounty{}

func init() {
	f := CDanzig但泽.(*但泽DanzigCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "368",
		Title:     "danzig",
		TitleName: "但泽",
		TitleCode: "c_danzig",
		Baronies:  map[string]feud.Barony{},
	}

	f.劳恩堡Danlauenburg = BDanlauenburg劳恩堡
	f.劳恩堡Danlauenburg.SetParent(f)
	
	f.但泽Danzig = BDanzig但泽
	f.但泽Danzig.SetParent(f)
	
	f.梅韦Mewe = BMewe梅韦
	f.梅韦Mewe.SetParent(f)
	
	f.奥利瓦Oliva = BOliva奥利瓦
	f.奥利瓦Oliva.SetParent(f)
	
	f.普茨克Puck = BPuck普茨克
	f.普茨克Puck.SetParent(f)
	
	f.施洛豪Schlochau = BSchlochau施洛豪
	f.施洛豪Schlochau.SetParent(f)
	
	f.施韦茨Schwetz = BSchwetz施韦茨
	f.施韦茨Schwetz.SetParent(f)
	
	f.图赫尔Tuchel = BTuchel图赫尔
	f.图赫尔Tuchel.SetParent(f)
	
}
