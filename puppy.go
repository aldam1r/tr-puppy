package trpuppy

import (
	dog "github.com/aldam1r/tr-dog"
)

func Bark() string {
	return "Woof"

}

func Barks() string {
	return "Woof, Woof, Woof"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func MoreBigBarks() string {
	return dog.WhenGrownUp(Barks()) + dog.WhenGrownUp(Barks())
}
