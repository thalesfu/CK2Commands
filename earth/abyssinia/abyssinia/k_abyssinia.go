package abyssinia

import (
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/afar"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/axum"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/damot"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/gojjam"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/gondar"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/harer"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/semien"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/shewa"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/wag"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AbyssiniaKingdom interface {
	feud.Kingdom
	DAfar阿法尔() afar.AfarDuke
	DAxum阿克苏姆() axum.AxumDuke
	DDamot达莫特() damot.DamotDuke
	DGojjam戈贾姆() gojjam.GojjamDuke
	DGondar贡德尔() gondar.GondarDuke
	DHarer拨拔力() harer.HarerDuke
	DSemien塞米恩() semien.SemienDuke
	DShewa绍阿() shewa.ShewaDuke
	DWag瓦格() wag.WagDuke
}

type 阿比西尼亚AbyssiniaKingdom struct {
	feud.BaseKingdom
	阿法尔Afar   afar.AfarDuke
	阿克苏姆Axum  axum.AxumDuke
	达莫特Damot  damot.DamotDuke
	戈贾姆Gojjam gojjam.GojjamDuke
	贡德尔Gondar gondar.GondarDuke
	拨拔力Harer  harer.HarerDuke
	塞米恩Semien semien.SemienDuke
	绍阿Shewa   shewa.ShewaDuke
	瓦格Wag     wag.WagDuke
}

func (k *阿比西尼亚AbyssiniaKingdom) DAfar阿法尔() afar.AfarDuke {
	return k.阿法尔Afar
}

func (k *阿比西尼亚AbyssiniaKingdom) DAxum阿克苏姆() axum.AxumDuke {
	return k.阿克苏姆Axum
}

func (k *阿比西尼亚AbyssiniaKingdom) DDamot达莫特() damot.DamotDuke {
	return k.达莫特Damot
}

func (k *阿比西尼亚AbyssiniaKingdom) DGojjam戈贾姆() gojjam.GojjamDuke {
	return k.戈贾姆Gojjam
}

func (k *阿比西尼亚AbyssiniaKingdom) DGondar贡德尔() gondar.GondarDuke {
	return k.贡德尔Gondar
}

func (k *阿比西尼亚AbyssiniaKingdom) DHarer拨拔力() harer.HarerDuke {
	return k.拨拔力Harer
}

func (k *阿比西尼亚AbyssiniaKingdom) DSemien塞米恩() semien.SemienDuke {
	return k.塞米恩Semien
}

func (k *阿比西尼亚AbyssiniaKingdom) DShewa绍阿() shewa.ShewaDuke {
	return k.绍阿Shewa
}

func (k *阿比西尼亚AbyssiniaKingdom) DWag瓦格() wag.WagDuke {
	return k.瓦格Wag
}

var KAbyssinia阿比西尼亚 AbyssiniaKingdom = &阿比西尼亚AbyssiniaKingdom{}

func init() {
	f := KAbyssinia阿比西尼亚.(*阿比西尼亚AbyssiniaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "abyssinia",
		TitleName: "阿比西尼亚",
		TitleCode: "k_abyssinia",
		Dukes:     map[string]feud.Duke{},
	}

	f.阿法尔Afar = afar.DAfar阿法尔
	f.阿法尔Afar.SetParent(f)

	f.阿克苏姆Axum = axum.DAxum阿克苏姆
	f.阿克苏姆Axum.SetParent(f)

	f.达莫特Damot = damot.DDamot达莫特
	f.达莫特Damot.SetParent(f)

	f.戈贾姆Gojjam = gojjam.DGojjam戈贾姆
	f.戈贾姆Gojjam.SetParent(f)

	f.贡德尔Gondar = gondar.DGondar贡德尔
	f.贡德尔Gondar.SetParent(f)

	f.拨拔力Harer = harer.DHarer拨拔力
	f.拨拔力Harer.SetParent(f)

	f.塞米恩Semien = semien.DSemien塞米恩
	f.塞米恩Semien.SetParent(f)

	f.绍阿Shewa = shewa.DShewa绍阿
	f.绍阿Shewa.SetParent(f)

	f.瓦格Wag = wag.DWag瓦格
	f.瓦格Wag.SetParent(f)

}
