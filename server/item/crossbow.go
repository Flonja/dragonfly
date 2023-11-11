package item

import (
	"fmt"
	"time"
	_ "unsafe"
)

// Crossbow is a ranged weapon similar to a bow that uses arrows or fireworks as ammunition.
type Crossbow struct {
	nopReleasable

	Item Stack
}

// MaxCount always returns 1.
func (Crossbow) MaxCount() int {
	return 1
}

// DurabilityInfo ...
func (Crossbow) DurabilityInfo() DurabilityInfo {
	return DurabilityInfo{
		MaxDurability: 464,
		BrokenItem:    simpleItem(Stack{}),
	}
}

// Release ...
func (c Crossbow) Release(releaser Releaser, duration time.Duration, ctx *UseContext) {
	creative := releaser.GameMode().CreativeInventory()
	if c.Item.Empty() {
		ticks := duration.Milliseconds() / 50
		if ticks < 25 {
			// The player must hold the crossbow for at least twenty-five ticks.
			return
		}

		projectileItem, ok := ctx.FirstFunc(func(stack Stack) bool {
			_, arrow := stack.Item().(Arrow)
			_, firework := stack.Item().(Firework)
			return arrow || firework
		})
		if !ok {
			if !creative {
				// No arrows/fireworks in inventory and not in creative mode.
				return
			}
			projectileItem = NewStack(Arrow{}, 1)
		}

		c.Item = projectileItem.Grow(-projectileItem.Count() + 1)
		if !creative {
			ctx.Consume(c.Item)
		}
		held, left := releaser.HeldItems()
		crossbow := newCrossbowWith(held, c)
		releaser.SetHeldItems(crossbow, left)
		fmt.Printf("%#v\n", writeItem(crossbow, true))

		return
	}

	c.Item = Stack{}
	held, left := releaser.HeldItems()
	releaser.SetHeldItems(newCrossbowWith(held, c), left)
	//d := float64(ticks) / 20
	//force := math.Min((d*d+d*2)/3, 1)
	//if force < 0.1 {
	//	// The force must be at least 0.1.
	//	return
	//}
	//
	//rot := releaser.Rotation()
	//rot = cube.Rotation{-rot[0], -rot[1]}
	//if rot[0] > 180 {
	//	rot[0] = 360 - rot[0]
	//}
	//var tip potion.Potion
	//if !arrow.Empty() {
	//	// Arrow is empty if not found in the creative inventory.
	//	tip = arrow.Item().(Arrow).Tip
	//}
	//
	//held, _ := releaser.HeldItems()
	//damage, consume := 2.0, 0, time.Duration(0), !creative
	//
	//create := releaser.World().EntityRegistry().Config().Arrow
	//projectile := create(eyePosition(releaser), releaser.Rotation().Vec3().Mul(force*5), rot, damage, releaser, force >= 1, false, !creative && consume, punchLevel, tip)
	//
	//ctx.DamageItem(1)
	//if consume {
	//	ctx.Consume(arrow.Grow(-arrow.Count() + 1))
	//}
	//
	//releaser.PlaySound(sound.BowShoot{})
	//releaser.World().AddEntity(projectile)
}

// FuelInfo ...
func (Crossbow) FuelInfo() FuelInfo {
	return newFuelInfo(time.Second * 10)
}

// EnchantmentValue ...
func (Crossbow) EnchantmentValue() int {
	return 1
}

// DecodeNBT ...
func (c Crossbow) DecodeNBT(data map[string]any) any {
	c.Item = mapItem(data, "chargedItem")
	return c
}

// EncodeNBT ...
func (c Crossbow) EncodeNBT() map[string]any {
	if !c.Item.Empty() {
		return map[string]any{
			"chargedItem": writeItem(c.Item, true),
		}
	}
	return nil
}

// EncodeItem ...
func (Crossbow) EncodeItem() (name string, meta int16) {
	return "minecraft:crossbow", 0
}

// newCrossbowWith duplicates an item.Stack with the new item type given.
func newCrossbowWith(input Stack, item Crossbow) Stack {
	if _, ok := input.Item().(Crossbow); !ok {
		return Stack{}
	}
	outputStack := NewStack(item, input.Count()).
		Damage(input.MaxDurability() - input.Durability()).
		WithCustomName(input.CustomName()).
		WithLore(input.Lore()...).
		WithEnchantments(input.Enchantments()...).
		WithAnvilCost(input.AnvilCost())
	for k, v := range input.Values() {
		outputStack = outputStack.WithValue(k, v)
	}
	return outputStack
}

// noinspection ALL
//
//go:linkname writeItem github.com/df-mc/dragonfly/server/internal/nbtconv.WriteItem
func writeItem(s Stack, disk bool) map[string]any

// noinspection ALL
//
//go:linkname mapItem github.com/df-mc/dragonfly/server/internal/nbtconv.MapItem
func mapItem(x map[string]any, k string) Stack
