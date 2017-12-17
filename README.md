[![Build Status](https://travis-ci.org/jacsmith21/repo.png?branch=master)](https://travis-ci.org/jacsmith21/repo)
# Repo
Use these datasets to quickly test your models.

## Usage
Get this repository:
```
$ go get github.com/jacsmith21/repo
```
then load the data:
```
bytes := repo.LoadBreastCancerWisconsinDataSet()
...
bytes := repo.LoadCarEvaluationDataSet()
...
bytes := repo.LoadEcoliDataSet()
...
bytes := repo.LoadLetterRecognitionDataSet()
...
bytes := repo.LoadMushroomDataSet()
```
