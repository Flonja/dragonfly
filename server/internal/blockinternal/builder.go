package blockinternal

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/block/customblock"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"strings"
)

// ComponentBuilder represents a builder that can be used to construct a block components map to be sent to a client.
type ComponentBuilder struct {
	permutations []map[string]any
	properties   []map[string]any
	components   map[string]any

	identifier string
	group      []world.CustomBlock
	traits     map[string][]any
}

// NewComponentBuilder returns a new component builder with the provided block data.
func NewComponentBuilder(identifier string, group []world.CustomBlock) *ComponentBuilder {
	traits := make(map[string][]any)
	for _, b := range group {
		_, properties := b.EncodeBlock()
		for trait, value := range properties {
			if _, ok := traits[trait]; !ok {
				traits[trait] = []any{}
			}
			traits[trait] = append(traits[trait], value)
		}
	}
	return &ComponentBuilder{
		properties: make([]map[string]any, 0),
		components: make(map[string]any),

		identifier: identifier,
		traits:     traits,
		group:      group,
	}
}

// AddProperty adds the provided property to the builder.
func (builder *ComponentBuilder) AddProperty(value map[string]any) {
	builder.properties = append(builder.properties, value)
}

// AddComponent adds the provided component to the builder.
func (builder *ComponentBuilder) AddComponent(name string, value any) {
	builder.components[name] = value
}

// AddPermutation adds a permutation to the builder.
func (builder *ComponentBuilder) AddPermutation(condition string, components map[string]any) {
	builder.permutations = append(builder.permutations, map[string]any{
		"condition":  condition,
		"components": components,
	})
}

// Trait finds a trait which satisfies all given values.
func (builder *ComponentBuilder) Trait(desired ...any) (string, bool) {
	for trait, values := range builder.traits {
		if len(values) != len(values) {
			// Not the same length, can't possibly be a match.
			continue
		}
		for i := range desired {
			if values[i] != desired[i] {
				continue
			}
		}
		return trait, true
	}
	return "", false
}

// Construct constructs the final block components map and returns it. It also applies the default properties required
// for the block to work without modifying the original maps in the builder.
func (builder *ComponentBuilder) Construct() map[string]any {
	permutations := slices.Clone(builder.permutations)
	properties := slices.Clone(builder.properties)
	components := maps.Clone(builder.components)
	builder.applyDefaultProperties(&properties)
	builder.applyDefaultComponents(components)
	result := map[string]any{"components": components}
	if len(properties) > 0 {
		result["properties"] = properties
	}
	if len(permutations) > 0 {
		result["molangVersion"] = int32(0)
		result["permutations"] = permutations
	}
	return result
}

// applyDefaultProperties applies the default properties to the provided map. It is important that this method does
// not modify the builder's properties map directly otherwise Empty() will return false in future use of the builder.
func (builder *ComponentBuilder) applyDefaultProperties(x *[]map[string]any) {
	for trait, values := range builder.traits {
		*x = append(*x, map[string]any{"enum": values, "name": trait})
	}
}

// applyDefaultComponents applies the default components to the provided map. It is important that this method does not
// modify the builder's components map directly otherwise Empty() will return false in future use of the builder.
func (builder *ComponentBuilder) applyDefaultComponents(x map[string]any) {
	base := builder.group[0]
	name := strings.Split(builder.identifier, ":")[1]
	materials := make(map[customblock.MaterialTarget]customblock.Material)
	textures, method := base.Textures()
	for target := range textures {
		materials[target] = customblock.NewMaterial(fmt.Sprintf("%v_%v", name, target.Name()), method)
	}

	origin := mgl64.Vec3{-8, 0, -8}
	size := mgl64.Vec3{16, 16, 16}
	geometries, ok := base.Geometries()
	if ok {
		geo := geometries.Geometry[0]
		origin, size = geo.Origin(), geo.Size()
	}

	model := customblock.NewModel(geometries.Geometry, origin, size)
	for target, material := range materials {
		model = model.WithMaterial(target, material)
	}
	for key, value := range model.Encode() {
		x[key] = value
	}
}