package repo

func LoadBreastCancerWisconsinDataSet() []byte {
  return load("breast-cancer-wisconsin.data")
}

func LoadCarEvaluationDataSet() []byte {
  return load("car.data")
}

func LoadEcoliDataSet() []byte {
  return load("ecoli.data")
}

func LoadLetterRecognitionDataSet() []byte {
  return load("letter-recognition.data")
}

func LoadMushroomDataSet() []byte {
  return load("mushroom.data")
}

func load(filename string) []byte {
  data, err := Asset("data/"+filename)
  if err != nil {
  	panic("asset not found")
  }

  return data
}
