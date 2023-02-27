package yuryev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YuryevCounty interface {
    feud.County
    BBavleny巴夫列内() 	feud.Barony
    BBerezhok别列若克() 	feud.Barony
    BGavrilov_posad加夫里洛夫镇() 	feud.Barony
    BGodenovo戈杰诺沃() 	feud.Barony
    BKolchugino科利丘吉诺() 	feud.Barony
    BVarvarino瓦尔瓦里诺() 	feud.Barony
    BYuryev尤里耶夫() 	feud.Barony
}

type 尤里耶夫YuryevCounty struct {
	feud.BaseCounty
	巴夫列内Bavleny 	feud.Barony
	别列若克Berezhok 	feud.Barony
	加夫里洛夫镇Gavrilov_posad 	feud.Barony
	戈杰诺沃Godenovo 	feud.Barony
	科利丘吉诺Kolchugino 	feud.Barony
	瓦尔瓦里诺Varvarino 	feud.Barony
	尤里耶夫Yuryev 	feud.Barony
}

func (c *尤里耶夫YuryevCounty) BBavleny巴夫列内() feud.Barony {
	return c.巴夫列内Bavleny
}
    
func (c *尤里耶夫YuryevCounty) BBerezhok别列若克() feud.Barony {
	return c.别列若克Berezhok
}
    
func (c *尤里耶夫YuryevCounty) BGavrilov_posad加夫里洛夫镇() feud.Barony {
	return c.加夫里洛夫镇Gavrilov_posad
}
    
func (c *尤里耶夫YuryevCounty) BGodenovo戈杰诺沃() feud.Barony {
	return c.戈杰诺沃Godenovo
}
    
func (c *尤里耶夫YuryevCounty) BKolchugino科利丘吉诺() feud.Barony {
	return c.科利丘吉诺Kolchugino
}
    
func (c *尤里耶夫YuryevCounty) BVarvarino瓦尔瓦里诺() feud.Barony {
	return c.瓦尔瓦里诺Varvarino
}
    
func (c *尤里耶夫YuryevCounty) BYuryev尤里耶夫() feud.Barony {
	return c.尤里耶夫Yuryev
}
    
var CYuryev尤里耶夫 YuryevCounty = &尤里耶夫YuryevCounty{}

func init() {
	f := CYuryev尤里耶夫.(*尤里耶夫YuryevCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1680",
		Title:     "yuryev",
		TitleName: "尤里耶夫",
		TitleCode: "c_yuryev",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴夫列内Bavleny = BBavleny巴夫列内
	f.巴夫列内Bavleny.SetParent(f)
	
	f.别列若克Berezhok = BBerezhok别列若克
	f.别列若克Berezhok.SetParent(f)
	
	f.加夫里洛夫镇Gavrilov_posad = BGavrilov_posad加夫里洛夫镇
	f.加夫里洛夫镇Gavrilov_posad.SetParent(f)
	
	f.戈杰诺沃Godenovo = BGodenovo戈杰诺沃
	f.戈杰诺沃Godenovo.SetParent(f)
	
	f.科利丘吉诺Kolchugino = BKolchugino科利丘吉诺
	f.科利丘吉诺Kolchugino.SetParent(f)
	
	f.瓦尔瓦里诺Varvarino = BVarvarino瓦尔瓦里诺
	f.瓦尔瓦里诺Varvarino.SetParent(f)
	
	f.尤里耶夫Yuryev = BYuryev尤里耶夫
	f.尤里耶夫Yuryev.SetParent(f)
	
}
