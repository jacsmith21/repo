package main

import "io/ioutil"

func LoadBreastCancerWisconsinDataSet() []byte {
  return loadBytes("breast-cancer-wisconsin.data")
}

func LoadCarEvaluationDataSet() []byte {
  return loadBytes("car.data")
}

func LoadEcoliDataSet() []byte {
  return loadBytes("ecoli.data")
}

func LoadLetterRecognitionDataSet() []byte {
  return loadBytes("letter-recognition.data")
}

func LoadMushroomDataSet() []byte {
  return loadBytes("mushroom.data")
}

func loadBytes(filename string) []byte {
  bytes, err := ioutil.ReadFile(filename) // just pass the file name
  if err != nil {
      panic(err)
  }

  return bytes
}
