package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestLoadBreastCancerWisconsinDataSet(t *testing.T) {
  bytes := loadBytes("breast-cancer-wisconsin.data")
  assert.Equal(t, uint8('1'), bytes[0])
}

func TestLoadCarEvaluationDataSet(t *testing.T) {
  bytes := loadBytes("car.data")
  assert.Equal(t, uint8('v'), bytes[0])
}

func TestLoadEcoliDataSet(t *testing.T) {
  bytes := loadBytes("ecoli.data")
  assert.Equal(t, uint8('A'), bytes[0])
}

func TestLoadLetterRecognitionDataSet(t *testing.T) {
  bytes := loadBytes("letter-recognition.data")
  assert.Equal(t, uint8('T'), bytes[0])
}

func TestLoadMushroomDataSet(t *testing.T) {
  bytes := loadBytes("mushroom.data")
  assert.Equal(t, uint8('p'), bytes[0])
}
