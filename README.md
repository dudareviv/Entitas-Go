# Entitas-Go

Entitas-GO is a fast Entity Component System Framework (ECS) Go 1.17 port of [Entitas v1.13.0 for C# and Unity.](https://github.com/sschmid/Entitas-CSharp)

# Code Generator

1) Install the library in your project

```
go get github.com/dudareviv/Entitas-Go
```

2) Create context file, with components.
```
//Important! The file should contain this line
//go:generate go run github.com/dudareviv/Entitas-Go
```

### Game.go
```golang
package main

//go:generate go run github.com/dudareviv/Entitas-Go

type Position struct {
	X, Y float64
}

type Direction struct {
	X, Y float64
}

type Speed int

type Health int

type Sprite struct {
	tag  string
	W, H int
}
```

3) Use
```
go generate
```

4) A folder 'Entitas' will be created in your project, please do not edit it

5) Edit the context file, add methods and components, and then use again `go generate`

# Examples

## Entity
```go
// main.go file
package main

import (
	ecs "myProject/Entitas"
)

func main() {
	contexts := ecs.CreateContexts()
    game := contexts.Game()

    // System registration
    systems := ecs.CreateSystemPool()
    systems.Add(&Translate{})
    systems.Add(&ReactiveTranslate{})

    // Create entity
    player := game.CreateEntity()

    // Add component
	player.AddPosition(10, 30)
	player.AddDirection(0, 0)
	player.AddSpeed(5)

    // Remove component
    player.RemoveSpeed()

    // Replace component
    player.ReplacePosition(30, 10)

    // On or Off component
    player.OffDirection()
    player.OnDirection()

    // Destroy entity
    player.Destroy()

    // GameLoop
    systems.Init(contexts)
    for true {
        systems.Execute()
        systems.Clean()
    }
    systems.Exit(contexts)
}
```

## System
```go
type Translate struct {
    group ecs.Group
}

func (s *Translate) Initer(contexts ecs.Contexts) {
	game := contexts.Game()
    s.group = game.Group(ecs.NewMatcher().AllOf(ecs.Position))
}

func (s *Translate) Executer() {
	for _, e := range s.group.GetEntities() {
        pos := e.GetPosition()
		e.ReplacePosition(pos.X + 10, pos.X + 10)
	}
}
```

## Reactive system
```go
type ReactiveTranslate struct {
}

func (s *ReactiveTranslate) Trigger(contexts ecs.Contexts) ecs.Collector {
	game := contexts.Game()
	return game.Collector(ecs.NewMatcher().AllOf(ecs.Position)).OnUpdate().OnAdd()
}

func (s *ReactiveTranslate) Filter(entity *ecs.Entity) bool {
	return entity.Has(ecs.Position)
}

func (s *ReactiveTranslate) Executer(entities []*ecs.Entity) {
	for _, e := range entities {
		pos := e.GetPosition()
		e.ReplacePosition(pos.X+10, pos.X+10)
	}
}
```