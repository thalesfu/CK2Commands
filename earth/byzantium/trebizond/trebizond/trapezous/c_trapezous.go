package trapezous

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TrapezousCounty interface {
    feud.County
    BAlucra阿卢杰拉() 	feud.Barony
    BDereli代雷利() 	feud.Barony
    BKelkit凯尔基特() 	feud.Barony
    BKoralla霍雷拉() 	feud.Barony
    BPaiperta佩佩尔塔() 	feud.Barony
    BRizaion里泽翁() 	feud.Barony
    BRizini里宗() 	feud.Barony
    BTrapezous特拉佩祖斯() 	feud.Barony
}

type 特拉佩祖斯TrapezousCounty struct {
	feud.BaseCounty
	阿卢杰拉Alucra 	feud.Barony
	代雷利Dereli 	feud.Barony
	凯尔基特Kelkit 	feud.Barony
	霍雷拉Koralla 	feud.Barony
	佩佩尔塔Paiperta 	feud.Barony
	里泽翁Rizaion 	feud.Barony
	里宗Rizini 	feud.Barony
	特拉佩祖斯Trapezous 	feud.Barony
}

func (c *特拉佩祖斯TrapezousCounty) BAlucra阿卢杰拉() feud.Barony {
	return c.阿卢杰拉Alucra
}
    
func (c *特拉佩祖斯TrapezousCounty) BDereli代雷利() feud.Barony {
	return c.代雷利Dereli
}
    
func (c *特拉佩祖斯TrapezousCounty) BKelkit凯尔基特() feud.Barony {
	return c.凯尔基特Kelkit
}
    
func (c *特拉佩祖斯TrapezousCounty) BKoralla霍雷拉() feud.Barony {
	return c.霍雷拉Koralla
}
    
func (c *特拉佩祖斯TrapezousCounty) BPaiperta佩佩尔塔() feud.Barony {
	return c.佩佩尔塔Paiperta
}
    
func (c *特拉佩祖斯TrapezousCounty) BRizaion里泽翁() feud.Barony {
	return c.里泽翁Rizaion
}
    
func (c *特拉佩祖斯TrapezousCounty) BRizini里宗() feud.Barony {
	return c.里宗Rizini
}
    
func (c *特拉佩祖斯TrapezousCounty) BTrapezous特拉佩祖斯() feud.Barony {
	return c.特拉佩祖斯Trapezous
}
    
var CTrapezous特拉佩祖斯 TrapezousCounty = &特拉佩祖斯TrapezousCounty{}

func init() {
	f := CTrapezous特拉佩祖斯.(*特拉佩祖斯TrapezousCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "678",
		Title:     "trapezous",
		TitleName: "特拉佩祖斯",
		TitleCode: "c_trapezous",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿卢杰拉Alucra = BAlucra阿卢杰拉
	f.阿卢杰拉Alucra.SetParent(f)
	
	f.代雷利Dereli = BDereli代雷利
	f.代雷利Dereli.SetParent(f)
	
	f.凯尔基特Kelkit = BKelkit凯尔基特
	f.凯尔基特Kelkit.SetParent(f)
	
	f.霍雷拉Koralla = BKoralla霍雷拉
	f.霍雷拉Koralla.SetParent(f)
	
	f.佩佩尔塔Paiperta = BPaiperta佩佩尔塔
	f.佩佩尔塔Paiperta.SetParent(f)
	
	f.里泽翁Rizaion = BRizaion里泽翁
	f.里泽翁Rizaion.SetParent(f)
	
	f.里宗Rizini = BRizini里宗
	f.里宗Rizini.SetParent(f)
	
	f.特拉佩祖斯Trapezous = BTrapezous特拉佩祖斯
	f.特拉佩祖斯Trapezous.SetParent(f)
	
}
