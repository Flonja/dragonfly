package item

import (
	"github.com/df-mc/dragonfly/server/item/potion"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/sound"
)

// noinspection SpellCheckingInspection
func init() {
	world.RegisterItem(AmethystShard{})
	world.RegisterItem(Apple{})
	world.RegisterItem(Arrow{})
	world.RegisterItem(BakedPotato{})
	world.RegisterItem(Beef{Cooked: true})
	world.RegisterItem(Beef{})
	world.RegisterItem(BeetrootSoup{})
	world.RegisterItem(Beetroot{})
	world.RegisterItem(BlazePowder{})
	world.RegisterItem(BlazeRod{})
	world.RegisterItem(BoneMeal{})
	world.RegisterItem(Bone{})
	world.RegisterItem(BookAndQuill{})
	world.RegisterItem(Book{})
	world.RegisterItem(BottleOfEnchanting{})
	world.RegisterItem(Bowl{})
	world.RegisterItem(Bow{})
	world.RegisterItem(Bread{})
	world.RegisterItem(Brick{})
	world.RegisterItem(Bucket{})
	world.RegisterItem(CarrotOnAStick{})
	world.RegisterItem(Charcoal{})
	world.RegisterItem(Chicken{Cooked: true})
	world.RegisterItem(Chicken{})
	world.RegisterItem(ClayBall{})
	world.RegisterItem(Clock{})
	world.RegisterItem(Coal{})
	world.RegisterItem(Cod{Cooked: true})
	world.RegisterItem(Cod{})
	world.RegisterItem(Compass{})
	world.RegisterItem(Cookie{})
	world.RegisterItem(CopperIngot{})
	world.RegisterItem(Crossbow{})
	world.RegisterItem(Diamond{})
	world.RegisterItem(DiscFragment{})
	world.RegisterItem(DragonBreath{})
	world.RegisterItem(DriedKelp{})
	world.RegisterItem(EchoShard{})
	world.RegisterItem(Egg{})
	world.RegisterItem(Elytra{})
	world.RegisterItem(Emerald{})
	world.RegisterItem(EnchantedApple{})
	world.RegisterItem(EnchantedBook{})
	world.RegisterItem(EnderPearl{})
	world.RegisterItem(Feather{})
	world.RegisterItem(FermentedSpiderEye{})
	world.RegisterItem(FireCharge{})
	world.RegisterItem(Firework{})
	world.RegisterItem(FlintAndSteel{})
	world.RegisterItem(Flint{})
	world.RegisterItem(GhastTear{})
	world.RegisterItem(GlassBottle{})
	world.RegisterItem(GlisteringMelonSlice{})
	world.RegisterItem(GlowstoneDust{})
	world.RegisterItem(GoldIngot{})
	world.RegisterItem(GoldNugget{})
	world.RegisterItem(GoldenApple{})
	world.RegisterItem(GoldenCarrot{})
	world.RegisterItem(Gunpowder{})
	world.RegisterItem(HeartOfTheSea{})
	world.RegisterItem(Honeycomb{})
	world.RegisterItem(InkSac{Glowing: true})
	world.RegisterItem(InkSac{})
	world.RegisterItem(IronIngot{})
	world.RegisterItem(IronNugget{})
	world.RegisterItem(LapisLazuli{})
	world.RegisterItem(Leather{})
	world.RegisterItem(MagmaCream{})
	world.RegisterItem(MelonSlice{})
	world.RegisterItem(MushroomStew{})
	world.RegisterItem(Mutton{Cooked: true})
	world.RegisterItem(Mutton{})
	world.RegisterItem(NautilusShell{})
	world.RegisterItem(NetherBrick{})
	world.RegisterItem(NetherQuartz{})
	world.RegisterItem(NetherStar{})
	world.RegisterItem(NetheriteIngot{})
	world.RegisterItem(NetheriteScrap{})
	world.RegisterItem(Paper{})
	world.RegisterItem(PhantomMembrane{})
	world.RegisterItem(PoisonousPotato{})
	world.RegisterItem(PoppedChorusFruit{})
	world.RegisterItem(Porkchop{Cooked: true})
	world.RegisterItem(Porkchop{})
	world.RegisterItem(PrismarineCrystals{})
	world.RegisterItem(PrismarineShard{})
	world.RegisterItem(Pufferfish{})
	world.RegisterItem(PumpkinPie{})
	world.RegisterItem(RabbitFoot{})
	world.RegisterItem(RabbitHide{})
	world.RegisterItem(RabbitStew{})
	world.RegisterItem(Rabbit{Cooked: true})
	world.RegisterItem(Rabbit{})
	world.RegisterItem(RawCopper{})
	world.RegisterItem(RawGold{})
	world.RegisterItem(RawIron{})
	world.RegisterItem(RecoveryCompass{})
	world.RegisterItem(RottenFlesh{})
	world.RegisterItem(Salmon{Cooked: true})
	world.RegisterItem(Salmon{})
	world.RegisterItem(Scute{})
	world.RegisterItem(Shears{})
	world.RegisterItem(ShulkerShell{})
	world.RegisterItem(Slimeball{})
	world.RegisterItem(Snowball{})
	world.RegisterItem(SpiderEye{})
	world.RegisterItem(Spyglass{})
	world.RegisterItem(Stick{})
	world.RegisterItem(Sugar{})
	world.RegisterItem(TropicalFish{})
	world.RegisterItem(TurtleShell{})
	world.RegisterItem(WarpedFungusOnAStick{})
	world.RegisterItem(Wheat{})
	world.RegisterItem(WrittenBook{})
	for _, t := range ArmourTiers() {
		world.RegisterItem(Helmet{Tier: t})
		world.RegisterItem(Chestplate{Tier: t})
		world.RegisterItem(Leggings{Tier: t})
		world.RegisterItem(Boots{Tier: t})
	}
	for _, pattern := range BannerPatterns() {
		world.RegisterItem(BannerPattern{Type: pattern})
	}
	for _, c := range Colours() {
		world.RegisterItem(Dye{Colour: c})
		world.RegisterItem(FireworkStar{FireworkExplosion: FireworkExplosion{Colour: c}})
	}
	for _, horn := range sound.GoatHorns() {
		world.RegisterItem(GoatHorn{Type: horn})
	}
	for i, p := range potion.All() {
		if i > 4 {
			world.RegisterItem(Arrow{Tip: p})
		}
		world.RegisterItem(LingeringPotion{Type: p})
		world.RegisterItem(SplashPotion{Type: p})
		world.RegisterItem(Potion{Type: p})
	}
	for _, t := range ToolTiers() {
		world.RegisterItem(Pickaxe{Tier: t})
		world.RegisterItem(Axe{Tier: t})
		world.RegisterItem(Shovel{Tier: t})
		world.RegisterItem(Sword{Tier: t})
		world.RegisterItem(Hoe{Tier: t})
	}
	for _, disc := range sound.MusicDiscs() {
		world.RegisterItem(MusicDisc{DiscType: disc})
	}
	for _, stew := range StewTypes() {
		world.RegisterItem(SuspiciousStew{Type: stew})
	}
	for _, sherd := range SherdTypes() {
		world.RegisterItem(PotterySherd{Type: sherd})
	}
}
