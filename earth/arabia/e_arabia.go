package arabia

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt"
	"github.com/thalesfu/CK2Commands/earth/arabia/israel"
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem"
	"github.com/thalesfu/CK2Commands/earth/arabia/syria"
	"github.com/thalesfu/CK2Commands/earth/arabia/yemen"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ArabiaEmpire interface {
    feud.Empire
    KArabia阿拉伯() 	arabia.ArabiaKingdom
    KEgypt埃及() 	egypt.EgyptKingdom
    KIsrael以色列() 	israel.IsraelKingdom
    KJerusalem耶路撒冷() 	jerusalem.JerusalemKingdom
    KSyria叙利亚() 	syria.SyriaKingdom
    KYemen也门() 	yemen.YemenKingdom
}

type 阿拉伯帝国ArabiaEmpire struct {
	feud.BaseEmpire
	阿拉伯Arabia 	arabia.ArabiaKingdom
	埃及Egypt 	egypt.EgyptKingdom
	以色列Israel 	israel.IsraelKingdom
	耶路撒冷Jerusalem 	jerusalem.JerusalemKingdom
	叙利亚Syria 	syria.SyriaKingdom
	也门Yemen 	yemen.YemenKingdom
}

func (e *阿拉伯帝国ArabiaEmpire) KArabia阿拉伯() arabia.ArabiaKingdom {
	return e.阿拉伯Arabia
}
    
func (e *阿拉伯帝国ArabiaEmpire) KEgypt埃及() egypt.EgyptKingdom {
	return e.埃及Egypt
}
    
func (e *阿拉伯帝国ArabiaEmpire) KIsrael以色列() israel.IsraelKingdom {
	return e.以色列Israel
}
    
func (e *阿拉伯帝国ArabiaEmpire) KJerusalem耶路撒冷() jerusalem.JerusalemKingdom {
	return e.耶路撒冷Jerusalem
}
    
func (e *阿拉伯帝国ArabiaEmpire) KSyria叙利亚() syria.SyriaKingdom {
	return e.叙利亚Syria
}
    
func (e *阿拉伯帝国ArabiaEmpire) KYemen也门() yemen.YemenKingdom {
	return e.也门Yemen
}
    
var EArabia阿拉伯帝国 ArabiaEmpire = &阿拉伯帝国ArabiaEmpire{}

func init() {
	f := EArabia阿拉伯帝国.(*阿拉伯帝国ArabiaEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "arabia",
		TitleName: "阿拉伯帝国",
		TitleCode: "e_arabia",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.阿拉伯Arabia = arabia.KArabia阿拉伯
	f.阿拉伯Arabia.SetParent(f)
	f.埃及Egypt = egypt.KEgypt埃及
	f.埃及Egypt.SetParent(f)
	f.以色列Israel = israel.KIsrael以色列
	f.以色列Israel.SetParent(f)
	f.耶路撒冷Jerusalem = jerusalem.KJerusalem耶路撒冷
	f.耶路撒冷Jerusalem.SetParent(f)
	f.叙利亚Syria = syria.KSyria叙利亚
	f.叙利亚Syria.SetParent(f)
	f.也门Yemen = yemen.KYemen也门
	f.也门Yemen.SetParent(f)
}
