package repo

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestLoadBreastCancerWisconsinDataSet(t *testing.T) {
  assert.NotPanics(t, func() { LoadBreastCancerWisconsinDataSet() })
}

func TestLoadCarEvaluationDataSet(t *testing.T) {
  assert.NotPanics(t, func() { LoadCarEvaluationDataSet() })
}

func TestLoadEcoliDataSet(t *testing.T) {
  assert.NotPanics(t, func() { LoadEcoliDataSet() })
}

func TestLoadLetterRecognitionDataSet(t *testing.T) {
  assert.NotPanics(t, func() { LoadLetterRecognitionDataSet() })
}

func TestLoadMushroomDataSet(t *testing.T) {
  assert.NotPanics(t, func() { LoadMushroomDataSet() })
}
