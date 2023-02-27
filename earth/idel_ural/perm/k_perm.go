package perm

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/bashkirs"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/komi"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/perm"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/ural"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/ustug"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/votyaki"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/zyriane"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PermKingdom interface {
    feud.Kingdom
    DBashkirs巴什基里亚() 	bashkirs.BashkirsDuke
    DKomi科米() 	komi.KomiDuke
    DPerm彼尔姆() 	perm.PermDuke
    DUral乌拉尔() 	ural.UralDuke
    DUstug大乌斯秋格() 	ustug.UstugDuke
    DVotyaki沃佳克() 	votyaki.VotyakiDuke
    DZyriane济良() 	zyriane.ZyrianeDuke
}

type 彼尔姆PermKingdom struct {
	feud.BaseKingdom
	巴什基里亚Bashkirs 	bashkirs.BashkirsDuke
	科米Komi 	komi.KomiDuke
	彼尔姆Perm 	perm.PermDuke
	乌拉尔Ural 	ural.UralDuke
	大乌斯秋格Ustug 	ustug.UstugDuke
	沃佳克Votyaki 	votyaki.VotyakiDuke
	济良Zyriane 	zyriane.ZyrianeDuke
}

func (k *彼尔姆PermKingdom) DBashkirs巴什基里亚() bashkirs.BashkirsDuke {
	return k.巴什基里亚Bashkirs
}
    
func (k *彼尔姆PermKingdom) DKomi科米() komi.KomiDuke {
	return k.科米Komi
}
    
func (k *彼尔姆PermKingdom) DPerm彼尔姆() perm.PermDuke {
	return k.彼尔姆Perm
}
    
func (k *彼尔姆PermKingdom) DUral乌拉尔() ural.UralDuke {
	return k.乌拉尔Ural
}
    
func (k *彼尔姆PermKingdom) DUstug大乌斯秋格() ustug.UstugDuke {
	return k.大乌斯秋格Ustug
}
    
func (k *彼尔姆PermKingdom) DVotyaki沃佳克() votyaki.VotyakiDuke {
	return k.沃佳克Votyaki
}
    
func (k *彼尔姆PermKingdom) DZyriane济良() zyriane.ZyrianeDuke {
	return k.济良Zyriane
}
    
var KPerm彼尔姆 PermKingdom = &彼尔姆PermKingdom{}

func init() {
	f := KPerm彼尔姆.(*彼尔姆PermKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "perm",
		TitleName: "彼尔姆",
		TitleCode: "k_perm",
		Dukes:  map[string]feud.Duke{},
	}

	f.巴什基里亚Bashkirs = bashkirs.DBashkirs巴什基里亚
	f.巴什基里亚Bashkirs.SetParent(f)
	
	f.科米Komi = komi.DKomi科米
	f.科米Komi.SetParent(f)
	
	f.彼尔姆Perm = perm.DPerm彼尔姆
	f.彼尔姆Perm.SetParent(f)
	
	f.乌拉尔Ural = ural.DUral乌拉尔
	f.乌拉尔Ural.SetParent(f)
	
	f.大乌斯秋格Ustug = ustug.DUstug大乌斯秋格
	f.大乌斯秋格Ustug.SetParent(f)
	
	f.沃佳克Votyaki = votyaki.DVotyaki沃佳克
	f.沃佳克Votyaki.SetParent(f)
	
	f.济良Zyriane = zyriane.DZyriane济良
	f.济良Zyriane.SetParent(f)
	
}
