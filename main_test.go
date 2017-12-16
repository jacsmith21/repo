package repo

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestLoadBreastCancerWisconsinDataSet(t *testing.T) {
  str := loadBytes("breast-cancer-wisconsin.data")
  assert.Equal(t, uint8('1'), str[0])
}

func TestLoadCarEvaluationDataSet(t *testing.T) {
  str := loadBytes("car.data")
  assert.Equal(t, uint8('v'), str[0])
}

func TestLoadEcoliDataSet(t *testing.T) {
  str := loadBytes("ecoli.data")
  assert.Equal(t, uint8('A'), str[0])
}

func TestLoadLetterRecognitionDataSet(t *testing.T) {
  str := loadBytes("letter-recognition.data")
  assert.Equal(t, uint8('T'), str[0])
}

func TestLoadMushroomDataSet(t *testing.T) {
  str := loadBytes("mushroom.data")
  assert.Equal(t, uint8('p'), str[0])
}
