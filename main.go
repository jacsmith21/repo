package repo

import "io/ioutil"

func LoadBreastCancerWisconsinDataSet() string {
  return string(loadBytes("breast-cancer-wisconsin.data"))
}

func LoadCarEvaluationDataSet() string {
  return string(loadBytes("car.data"))
}

func LoadEcoliDataSet() string {
  return string(loadBytes("ecoli.data"))
}

func LoadLetterRecognitionDataSet() string {
  return string(loadBytes("letter-recognition.data"))
}

func LoadMushroomDataSet() string {
  return string(loadBytes("mushroom.data"))
}

func loadBytes(filename string) []byte {
  bytes, err := ioutil.ReadFile(filename) // just pass the file name
  if err != nil {
      panic(err)
  }

  return bytes
}
