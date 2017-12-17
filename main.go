package repo

import "os"

func LoadBreastCancerWisconsinDataSet() *os.File {
  return open("breast-cancer-wisconsin.data")
}

func LoadCarEvaluationDataSet() *os.File {
  return open("car.data")
}

func LoadEcoliDataSet() *os.File {
  return open("ecoli.data")
}

func LoadLetterRecognitionDataSet() *os.File {
  return open("letter-recognition.data")
}

func LoadMushroomDataSet() *os.File {
  return open("mushroom.data")
}

func open(filename string) *os.File {
  file, err := os.Open(filename) // just pass the file name
  if err != nil {
      panic(err)
  }

  return file
}
