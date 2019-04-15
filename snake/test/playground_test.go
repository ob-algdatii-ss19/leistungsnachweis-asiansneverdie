package test

import (
	"github.com/Plateria/AlgoDatII/leistungsnachweis-asiansneverdie/snake/model"
	"testing"
)

func TestPlaygroundCreate(t *testing.T) {
	playground := model.NewPlayground()
	playground.Create()
}
