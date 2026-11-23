package spentenergy

import (
	"errors"
	"math"
	"time"
)

const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	switch {
	case steps <= 0:
		return 0, errors.New("количество шагов должно быть больше нуля")
	case weight <= 0 || math.IsNaN(weight) || math.IsInf(weight, 0):
		return 0, errors.New("некорректный вес")
	case height <= 0 || math.IsNaN(height) || math.IsInf(height, 0):
		return 0, errors.New("некорректный рост")
	case duration <= 0:
		return 0, errors.New("длительность должна быть больше нуля")
	}

	speed := MeanSpeed(steps, height, duration)
	return weight * speed * duration.Minutes() / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}